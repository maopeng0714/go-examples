package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

// image dir
const (
	UploadDir = "./uploads"
	ViewDir   = "./views"
	ListDir   = 0x0001
)

func check(err error) {
	if err != nil {
		panic(err)

	}
}
func renderHTML(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {

	t := templates[tmpl+".html"]
	e := t.Execute(w, locals)
	check(e)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		renderHTML(w, "upload", nil)
	}
	if r.Method == "POST" {
		f, h, e := r.FormFile("image")
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		savefile, e := os.Create(UploadDir + "/" + filename)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		defer savefile.Close()
		if _, e := io.Copy(savefile, f); e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := UploadDir + "/" + imageID

	if ie := isExist(imagePath); !ie {
		http.NotFound(w, r)
		return
	}
	image, e := ioutil.ReadFile(imagePath)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		imme := http.DetectContentType(image)
		log.Printf("file type is: " + imme)
		w.Header().Set("Content-Type", imme)
		http.ServeFile(w, r, imagePath)
	}
}

func isExist(path string) bool {
	_, e := os.Stat(path)
	if e == nil {
		return true
	}
	return os.IsExist(e)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileinfors, e := ioutil.ReadDir(UploadDir)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
	images := []string{}
	imageMap := make(map[string]interface{})
	for _, fileinfo := range fileinfors {
		imageid := fileinfo.Name()
		images = append(images, imageid)
	}
	imageMap["images"] = images
	renderHTML(w, "list", imageMap)
}

var templates map[string]*template.Template

func init() {
	fileInfos, e := ioutil.ReadDir(ViewDir)
	check(e)
	templates = make(map[string]*template.Template)
	for _, fileinfo := range fileInfos {
		filename := fileinfo.Name()
		if ext := path.Ext(filename); ext != ".html" {
			continue
		}
		filepath := ViewDir + "/" + filename
		t := template.Must(template.ParseFiles(filepath))
		templates[filename] = t
	}
}

func safeHandler(handler http.HandlerFunc) (safeHandler http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", handler, e)
				log.Println(string(debug.Stack()))
			}
		}()
		handler(w, r)
	}
}

func staticHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if flags&ListDir == 0 {
			if exists := isExist(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}
func main() {
	mux := http.NewServeMux()
	staticHandler(mux, "/assets/", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

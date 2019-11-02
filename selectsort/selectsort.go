
package main

import "fmt"

func selectSort (a []int)  {
	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func main() {
	a := []int {2, 8, 1, 6, 5, 2, 7, 3, 4}
	if a == nil {
		return
	}
	
    if(len(a) == 1){
		return
	}
	selectSort(a)
	fmt.Println(a)
}
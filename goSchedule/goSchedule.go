package main

import (
	"fmt"
	"runtime"
)

func printOdd(n int, done *chan bool) {

	for i := 1; i <= n; i += 2 {
		fmt.Println(i)
		runtime.Gosched()
	}
	*done <- true
}

func printDouble(n int, done *chan bool) {

	for i := 2; i <= n; i += 2 {
		fmt.Println(i)
		runtime.Gosched()
	}
	*done <- true
}

func main() {

	done := make(chan bool, 2)
	n := 100

	go printOdd(n, &done)
	go printDouble(n, &done)
	defer close(done)
	<-done
	<-done
}

package main

import "fmt"

func printOdd(n int, oddCh, doubleCh, done *chan bool) {

	for i := 1; i <= n; i += 2 {
		<-*oddCh
		fmt.Println(i)
		*doubleCh <- true
	}
	*done <- true
}

func printDouble(n int, oddCh, doubleCh, done *chan bool) {

	for i := 2; i <= n; i += 2 {
		<-*doubleCh
		fmt.Println(i)
		*oddCh <- true
	}
	*done <- true
}

func main() {
	oddCh := make(chan bool, 1)
	doubleCh := make(chan bool, 1)
	done := make(chan bool, 2)
	n := 100
	oddCh <- true
	go printOdd(n, &oddCh, &doubleCh, &done)
	go printDouble(n, &oddCh, &doubleCh, &done)
	defer close(oddCh)
	defer close(doubleCh)
	defer close(done)
	<-done
	<-done
}

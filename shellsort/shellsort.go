package main

import "fmt"

func shellSort (a []int)  {
	for step := len(a)/2; step > 0; step /= 2 {
		for i := step; i < len(a); i++ {
			value := a[i]
			j := i - step
			for ; j  >= 0 && a[j] > value; j -= step {
				 a[j + step] = a[j]
			}
			a[j + step] = value
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
	shellSort(a)
	fmt.Println(a)
}
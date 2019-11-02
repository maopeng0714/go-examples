package main

import "fmt"

func mergeSort (a []int)  {
	
	for i := 1; i < len(a); i++ {
		value := a[i]
		j := i - 1
		for ; j  >= 0 && a[j] > value; j-- {
				a[j], a[j+1] = a[j+1], a[j]
		}
		a[j+1] = value		
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
	mergeSort(a)
	fmt.Println(a)
}
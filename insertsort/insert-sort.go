package main

import "fmt"

func insertSort (a []int)  {
	
	for i := 1; i < len(a) ; i++ {
		value := a[i]
		j := i - 1	
		for ; j  >= 0 ; j-- {
			if a[j] < value {
				break
			}
			a[j+1] = a[j]
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
	insertSort(a)
	fmt.Println(a)
}
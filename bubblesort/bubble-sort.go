package main

import "fmt"

func bubbleSort (a []int)  {
	
	for i := 0; i < len(a) -1; i++ {
		swap := false
		for j := 0; j < len(a) - i - 1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swap = true
			}
		}
		if !swap {
			return
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
	bubbleSort(a)
	fmt.Println(a)
}
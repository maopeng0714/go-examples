package main

import "fmt"

func updateHeap (a []int,  i, n int)  {
	max := i
	left := 2 * i + 1
	right := 2 * (i+1)

	if left < n && a[left] > a[max] {
		max = left
	}
	if right < n && a[right] > a[max] {
		max = right
	}
	if max != i {
		a[max], a[i] = a[i], a[max]
		updateHeap(a, max, n)
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
////这里元素的索引是从0开始的,所以最后一个非叶子结点array.length/2 - 1
    length := len(a)

	for i := length / 2 - 1; i >= 0; i-- {
		updateHeap(a, i, length)
	}

	for j := length - 1; j >= 0 ; j-- {
		a[j], a[0] = a[0], a[j]
		updateHeap(a, 0, j)
	}
	fmt.Println(a)
}
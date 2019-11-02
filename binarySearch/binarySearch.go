package main

import (
	"fmt"

	"github.com/maopeng0714/go-examples/quicksort"
)

func binarySearch(a []int, key, start, end int) int {
	mid := start + (key-a[start])/(a[end]-a[start])*(end-start)
	if key == a[mid] {
		return mid
	} else if key > a[mid] {
		return binarySearch(a, key, mid+1, end)
	} else {
		return binarySearch(a, key, start, mid-1)
	}
}
func main() {
	a := []int{2, 8, 1, 6, 5, 2, 7, 3, 4}
	if a == nil {
		return
	}

	if len(a) == 1 {
		return
	}
	////这里元素的索引是从0开始的,所以最后一个非叶子结点array.length/2 - 1
	length := len(a)
	quicksort.QuickSort(a, 0, length-1)
	fmt.Println("after sort the array is  ", a)
	pos := binarySearch(a, 5, 0, length-1)
	fmt.Println("find the position of 5 in the sorted array is ", pos)
}

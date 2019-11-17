package main

import (
	"fmt"
)

func bestProfit(p []int) int {
	n := len(p)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		max := -1
		for j := 0; j < i; j++ {
			value := p[j] + r[i-j]
			if max < value {
				max = value
			}
		}
		r[i] = max
	}
	return r[n-1]
}

func main() {

	p := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	profit := bestProfit(p)
	fmt.Println("the best profit is ", profit)

}

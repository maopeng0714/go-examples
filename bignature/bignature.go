package main

import (
	"fmt"
	"strconv"
)

func bigNature(a []int) int {
	result := ""
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			i_str := strconv.Itoa(a[i])
			j_str := strconv.Itoa(a[j])
			pre_str := i_str + j_str
			last_str := j_str + i_str
			pre_int, err1 := strconv.Atoi(pre_str)
			last_int, err2 := strconv.Atoi(last_str)
			if err1 == nil && err2 == nil && pre_int < last_int {
				a[i], a[j] = a[j], a[i]
			}
		}

		result += strconv.Itoa(a[i])
	}
	nature, _ := strconv.Atoi(result)
	return nature
}

func main() {
	a := []int{1, 3, 12, 31, 15, 37}
	result := bigNature(a)
	fmt.Println("The largest value is ", result)
}

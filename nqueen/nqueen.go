package main

import (
	"fmt"
	"math"
)

//http://blog.sina.com.cn/s/blog_696187fd0100p5ri.html
// 判断两个皇后是否在一条斜线或者对角线的条件是:
// 假设两个皇后的坐标分别是X1, Y1和X2，Y2；
// 则条件是|X1 - Y1| == |X2-Y2| || |X1 + Y1| == |X2 + Y2|
// 同理判断俩皇后在同一列的条件是Y1==Y2

//前提假设任意俩皇后肯定不在同一行上，X[n]是n个皇后所在列的坐标数组。

func place(k int, a []int) bool {
	for i := 0; i < k; i++ {

		if a[i] == a[k] || math.Abs(float64(k-a[k])) == math.Abs(float64(i-a[i])) {
			return false
		}
	}
	return true
}

func queen(t int, number *int, a *[]int) int {
	n := len(*a)
	if t >= n {
		*number++
	} else {
		for i := 0; i < n; i++ {
			(*a)[t] = i
			if place(t, *a) {
				queen(t+1, number, a)
			}
		}
	}
	return *number
}

func main() {
	//皇后数
	n := 4
	//皇后的列坐标数组
	a := make([]int, n)
	//解决方案数目
	number := 0
	number = queen(0, &number, &a)
	fmt.Println(number)
	fmt.Println(a)
}

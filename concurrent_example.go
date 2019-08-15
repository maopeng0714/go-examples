package main

import "fmt"

//Vector the float64 array
type Vector []float64

//DoSome sum the some part of the array
func (v Vector) DoSome(i, n int, c chan float64) {
	var sum float64
	fmt.Println("the start:", i, " and the end is ", n)
	for ; i < n; i++ {
		sum += v[i]
	}
	c <- sum
}

//NCPU the number of CPUs
const NCPU = 4

//DoAll sum the whole array
func (v Vector) DoAll() float64 {
	c := make(chan float64, NCPU)
	length := len(v)
	for index := 0; index < NCPU; index++ {
		go v.DoSome(index*length/NCPU, (index+1)*length/NCPU, c)
	}
	var sum float64
	for i := 0; i < NCPU; i++ {
		seq, ok := <-c
		if ok {
			fmt.Println("the seq:", i, " of the float array is ", seq)
			sum += seq
		}
	}
	return sum
}

func main() {
	var v Vector = []float64{1.1, 2.1, 3.1, 4.1, 6.1, 8.1, 9.1, 10.1, 11}
	var sum = v.DoAll()
	fmt.Println("Sum of the float array is ", sum)
	var expected float64
	for _, value := range v {
		expected += value
	}
	if sum == expected {
		fmt.Println("Expected: ", expected, "not equals Sum: ", sum)
	} else {
		fmt.Println("Expected: ", expected, "equals Sum: ", sum)
	}
}

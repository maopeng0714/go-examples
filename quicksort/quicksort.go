package quicksort

import "fmt"

//QuickSort the export function
func QuickSort(a []int, left, right int) {

	temp := a[left]
	p := left
	i, j := left, right
	//一趟操作的最终目的是把“支点”放到它应该去的地方
	for i <= j {
		for p <= j && a[j] >= temp {
			j--
		}
		if p < j {
			a[p] = a[j]
			p = j
		}
		for p >= i && a[i] <= temp {
			i++
		}
		if p > i {
			a[p] = a[i]
			p = i
		}
	}
	a[p] = temp

	if p-left > 1 {
		QuickSort(a, left, p-1)
	}
	if right-p > 1 {
		QuickSort(a, p+1, right)
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
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

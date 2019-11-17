package dp

func fib1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

func Fib2(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	mem0 := 0
	mem1 := 1
	mem2 := 1
	for i := 2; i <= n; i++ {
		mem2 = mem1 + mem0
		mem0 = mem1
		mem1 = mem2
	}
	return mem2
}

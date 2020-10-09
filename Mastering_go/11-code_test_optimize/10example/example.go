package examp

func fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

func length(s string) int {
	return len(s)
}

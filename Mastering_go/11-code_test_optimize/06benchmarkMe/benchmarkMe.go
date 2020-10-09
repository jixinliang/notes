package main

import "fmt"

func fib1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib1(n-1) + fib1(n-2)
	}
}

// this is a relatively simple and somewhat slow approach
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

func fib3(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var tmp int
		if i <= 2 {
			tmp = 1
		} else {
			tmp = fn[i-1] + fn[i-2]
		}
		fn[i] = tmp
	}
	return fn[n]
}

func main() {
	fmt.Println(fib1(40))
	fmt.Println(fib2(40))
	fmt.Println(fib3(40))
}

// go test -bench=. benchmarkMe*
// go test -benchmem -bench=. benchmarkMe*
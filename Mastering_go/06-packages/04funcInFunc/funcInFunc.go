package main

import "fmt"

func add(x int) int {
	return x + x
}

func multi(x int) int {
	return x * x
}

func funcInFunc(fn func(int) int, val int) int {
	return fn(val)
}

func main() {
	fmt.Println(add(10))
	fmt.Println(multi(10))
	fmt.Println(funcInFunc(func(i int) int {
		return i + i
	}, 10))
}

package main

import "fmt"

func inputPtr(x *float64) float64 {
	return *x * *x
}

func retPtr(x int) *int {
	y := x * x
	return &y
}

func main() {
	x := 12.2
	fmt.Println(inputPtr(&x))

	x = 12
	fmt.Println(inputPtr(&x))

	res := retPtr(10)
	fmt.Println("ptr of res:", res)
	fmt.Println("value of res:", *res)
}

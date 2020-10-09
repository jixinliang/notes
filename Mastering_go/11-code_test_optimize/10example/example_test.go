package examp

import "fmt"

func Examplefib1()  {
	fmt.Println(fib1(10))
	fmt.Println(fib1(2))
	// Output:
	// 55
	// 1
}

func Examplelength()  {
	fmt.Println(length("123456789"))
	fmt.Println(length(""))
	// Output:
	// 10
	// 0
}
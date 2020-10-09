package main

import "fmt"

func f1() int {
	fmt.Println("Calling f1()")
	return 1
	//return 10
}

func f2() int {
	if true {
		return 2
	}
	fmt.Println("Calling f2()")
	return 0
}

func main() {
	fmt.Println(f1())
	fmt.Println("Exiting")
}

// go vet unreachable.go
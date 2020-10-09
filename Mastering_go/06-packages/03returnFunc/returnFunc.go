package main

import "fmt"

func f1() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {
	ff := f1()
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())
}

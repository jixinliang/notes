package main

import "fmt"

func add(a int) int  {
	return a+a
}

func main() {
	i := 6
	fmt.Println("Debugging ...")
	fmt.Println(i)
	fmt.Println(add(i))
}

package main

import (
	"fmt"
	"os"
)

func variadicFunc(inputs ...string) {
	fmt.Println(inputs)
}

func oneByOne(msg string, s ...int) int {
	fmt.Println(msg)
	sum := 0
	for k, v := range s {
		fmt.Println(k, v)
		sum += v
	}
	s[0] = -100
	return sum
}

func main() {
	args := os.Args
	variadicFunc(args...)

	sum := oneByOne("Adding nums:", 1, 2, 3, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)

	s := []int{1, 2, 3}
	sum = oneByOne("Adding nums again:", s...)
	fmt.Println("s:", s)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("The program need one argument!")
		return
	}

	a, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	square := func(x int) int { return x * x }
	fmt.Println("The square of", a, "is", square(a))

	double := func(x int) int { return x + x }
	fmt.Println("The double of", a, "is", double(a))

	fmt.Println(doubleSquare(a))

	i, i2 := doubleSquare(a)
	fmt.Println(i, i2)
}

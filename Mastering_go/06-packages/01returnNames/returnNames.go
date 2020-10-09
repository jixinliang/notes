package main

import (
	"fmt"
	"os"
	"strconv"
)

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func namedMinMax1(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("The program need at least three arguments!")
		return
	}
	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[2])

	fmt.Println(namedMinMax(x, y))
	fmt.Println(namedMinMax1(x, y))
}

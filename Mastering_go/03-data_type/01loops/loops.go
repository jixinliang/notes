package main

import (
	"fmt"
)

func conBrake() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}
}

func while() {
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
}

func doWhile() {
	i := 0
	flag := true
	for ok := true; ok; ok = flag {
		if i > 10 {
			flag = false
		}
		fmt.Print(i, " ")
		i++
	}
}

func forRange() {
	myArr := [...]int{0, 1, -1, 2, -2}
	for idx, val := range myArr {
		fmt.Println("idx:", idx, "val:", val)
	}
}

func main() {
	forRange()
}

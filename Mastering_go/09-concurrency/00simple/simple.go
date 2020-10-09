package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func main() {
	go f1()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	time.Sleep(time.Second)
	fmt.Println()
}

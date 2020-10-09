package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second)
	duration := time.Since(start)
	fmt.Println("It took",duration,"to finish")

	start = time.Now()
	for i := 0; i < 200000000; i++ {
		_ = i
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop",duration,"to finish")

	sum := 0
	start = time.Now()
	for i := 0; i < 200000000; i++ {
		sum += i
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop",duration,"to finish")
}

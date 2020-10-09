package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func f1(cc chan chan int, b chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum += i
		}
		c <- sum
	case <-b:
		return
	}
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <number>\n", filepath.Base(args[0]))
		return
	}

	var times int
	times, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Atoi:", err)
		return
	}

	cc := make(chan chan int)

	for i := 1; i < times+1; i++ {
		b := make(chan bool)

		go f1(cc, b)

		tmp := <-cc
		tmp <- i
		for sum := range tmp {
			fmt.Println("Sum:", i, "->", sum)
		}
		time.Sleep(time.Second)
		close(b)
	}

}

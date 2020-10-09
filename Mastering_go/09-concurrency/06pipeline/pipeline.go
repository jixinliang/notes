package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var closed = false
var data = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if closed {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(in <-chan int, out chan<- int) {
	for x := range in {
		fmt.Print(x, " ")
		// check if the data aleardy existed in the map
		// then close the out chan
		_, ok := data[x]
		if ok {
			closed = true
		} else {
			data[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	fmt.Println("The sum of the random numbers is:", sum)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: go run %s [num1, num2]\n", filepath.Base(args[0]))
		return
	}

	n1, _ := strconv.Atoi(args[1])
	n2, _ := strconv.Atoi(args[2])

	rand.Seed(time.Now().Unix())

	out := make(chan int)
	in := make(chan int)
	go first(n1, n2, out)
	go second(out, in)
	third(in)
}

// go run -race .
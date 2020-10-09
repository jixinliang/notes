package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var readVal = make(chan int)
var writeVal = make(chan int)

func set(newVal int) {
	writeVal <- newVal
}

// get the last value of chan
func get() int {
	return <-readVal
}

func monitor() {
	var value int
	for {
		select {
		case newVal := <-writeVal:
			value = newVal
			// print random of numbers
			fmt.Printf("%d ",value)
		case readVal <- value: // send the last value to readVal chan cause unbuffer chan
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <Number>\n", filepath.Base(args[0]))
		return
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Atoi:", err)
		return
	}

	fmt.Printf("Going to create %d random numbers.\n", n)

	rand.Seed(time.Now().Unix())

	go monitor()

	var wg sync.WaitGroup

	// the key to set nums of random numbers
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}

	wg.Wait()
	fmt.Printf("\nLast value: %d\n", get())
}

// sharing by communicating

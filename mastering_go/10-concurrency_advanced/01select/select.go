package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// select statement for channel
func gen(min, max int, createNum chan int, end chan bool) {
	for {
		select {
		case createNum <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(time.Second * 2):
			fmt.Println("\nTime after!")
		}
	}
}

func main() {
	args := os.Args
	rand.Seed(time.Now().Unix())
	createNum := make(chan int)
	end := make(chan bool)

	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <number>\n", filepath.Base(args[0]))
		return
	}

	n, _ := strconv.Atoi(args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	go gen(0, 2*n, createNum, end)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNum)
	}

	time.Sleep(time.Second * 3)
	fmt.Println("Exiting...")
	end <- true
}

/*
If none of the channels in a select statement are ready, the select statement will block
until one of the channels is ready. If multiple channels of a select statement are ready,
then the Go runtime will make a random selection from the set of these ready channels
*/


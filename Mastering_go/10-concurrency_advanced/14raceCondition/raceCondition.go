package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

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

	var wg sync.WaitGroup
	// race condition variable for map key and for loop
	var i int

	k := make(map[int]int)
	k[1] = 12

	// the idea to address the problem is that reseting the variable i as pram
	// and add lock to limit the goroutine when write data

	var lock sync.RWMutex

	for i = 0; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			lock.Lock()
			k[j] = j
			lock.Unlock()
		}(i)
	}

	wg.Wait() // that position should be before the assignment then race
	k[2] = 10
	fmt.Printf("%+v\n", k)
}

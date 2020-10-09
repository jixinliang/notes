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


	for i = 0; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			k[j] = j
		}(i)
	}

	k[2] = 10
	wg.Wait()
	fmt.Printf("%+v\n", k)
}

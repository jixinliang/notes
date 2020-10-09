package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n",20,"Number of goroutines")
	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d goroutines\n", count)

	var wg sync.WaitGroup
	fmt.Printf("WaitGroup before: %#v\n", wg)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			// waiting for the for loop to executing
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("WaitGroup after: %#v\n", wg)

	wg.Wait()
	fmt.Println("\nExiting...")
}

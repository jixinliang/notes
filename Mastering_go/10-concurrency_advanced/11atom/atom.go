package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	Value int64
}

func (a *atomCounter) value() int64 {
	return atomic.LoadInt64(&a.Value)
}

func main() {
	ix := flag.Int("x", 100, "goroutine numbers")
	iy := flag.Int("y", 200, "value")
	flag.Parse()

	x := *ix
	y := *iy

	var wg sync.WaitGroup
	counter := atomCounter{}

	for i := 0; i < x; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			for i := 0; i < y; i++ {
				atomic.AddInt64(&counter.Value, 1)

				// make the program thread unsafe
				//counter.Value++
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Value:",counter.Value)
}

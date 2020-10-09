package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var (
	lock sync.Mutex
	v1   int
)

func change(i int) {
	lock.Lock()
	time.Sleep(time.Second)
	v1 += 1
	if v1%10 == 0 {
		v1 -= 10 * i
	}
	lock.Unlock()
}

func read() int {
	lock.Lock()
	res := v1
	lock.Unlock()
	return res
}

func main() {
	args := os.Args
	var wg sync.WaitGroup
	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <number>\n", filepath.Base(args[0]))
		return
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Atoi:",err)
		return
	}

	fmt.Printf("%d ", read())

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf(" -> %d", read())
		}(i)
	}

	wg.Wait()
	fmt.Printf("-> %d\n", read())
}

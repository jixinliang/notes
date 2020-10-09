package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(wg *sync.WaitGroup, dur time.Duration) bool {
	tmp := make(chan int)
	go func() {
		defer close(tmp)
		time.Sleep(time.Second * 3)

		wg.Wait()
	}()

	select {
	case <-tmp:
		return false
	case <-time.After(dur):
		return true
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Need a time duration!")
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	t, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Atoi:", err)
		return
	}

	dur := time.Duration(uint32(t)) * time.Millisecond
	fmt.Println("Timeout period is:", dur)

	if timeout(&wg, dur) {
		fmt.Println("Time out!")
	} else {
		fmt.Println("OK!")
	}

	wg.Done()

	if timeout(&wg, dur) {
		fmt.Println("Time out!")
	} else {
		fmt.Println("OK!")
	}
}

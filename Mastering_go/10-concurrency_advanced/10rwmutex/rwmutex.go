package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	rwLock sync.RWMutex
	lock   sync.Mutex
	passwd string
}

var passwd = secret{passwd: "666"}

func change(s *secret, passwd string) {
	s.rwLock.Lock()
	fmt.Println("Lock change")
	time.Sleep(time.Second * 6)
	s.passwd = passwd
	s.rwLock.Unlock()
}

func show(s *secret) string {
	s.rwLock.RLock()
	fmt.Println()
	time.Sleep(time.Second * 3)
	defer s.rwLock.RUnlock()
	return s.passwd
}

func showWithLock(s *secret) string {
	s.rwLock.Lock()
	fmt.Println("Show with lock")
	time.Sleep(time.Second * 3)
	defer s.rwLock.Unlock()
	return s.passwd
}

func main() {
	args := os.Args
	var showFunc = func(s *secret) string { return "" }

	if len(args) != 2 {
		fmt.Println("Using rwMutex!")
		showFunc = show
	} else {
		fmt.Println("Using mutex!")
		showFunc = showWithLock
	}

	var wg sync.WaitGroup
	fmt.Println("Password:", showFunc(&passwd))

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go passwd:", showFunc(&passwd))
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		change(&passwd, "123")
	}()

	wg.Wait()
	fmt.Println("Final Pass:", showFunc(&passwd))
}

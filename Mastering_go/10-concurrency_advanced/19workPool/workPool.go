package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

// input of the request
type Client struct {
	ID      int
	Integer int
}

// result of request
type Data struct {
	Job    Client
	Square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done() // not use the copy of wg

	for c := range clients {
		square := c.Integer * c.Integer
		output := Data{c, square}

		data <- output
		time.Sleep(time.Second)
	}
}

func makeWorkerPool(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(data)
}

func createJobs(n int) {
	for i := 1; i <= n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: go run %s <Jobs num> <Worker num>\n", filepath.Base(args[0]))
		return
	}

	jobs, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Atoi jobs:", err)
		return
	}

	workers, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Atoi worker:", err)
		return
	}

	go createJobs(jobs)

	finished := make(chan interface{})
	go func() {
		for data := range data {
			fmt.Printf("Client ID: %d \t", data.Job.ID)
			fmt.Println("Integer:", data.Job.Integer, "Square:", data.Square)
		}
		finished <- true
	}()

	makeWorkerPool(workers)
	fmt.Println("Finished: ", <-finished)

}

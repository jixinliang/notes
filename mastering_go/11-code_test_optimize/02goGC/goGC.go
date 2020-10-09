package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func getStats(memStats runtime.MemStats) {
	runtime.ReadMemStats(&memStats)
	fmt.Println("Alloc:", memStats.Alloc)
	fmt.Println("TotalAlloc:", memStats.TotalAlloc)
	fmt.Println("HeapAlloc:", memStats.HeapAlloc)
	fmt.Println("NumGC:", memStats.NumGC)
	fmt.Println("------------------")
}

func main() {
	file, err := os.Create("/tmp/traceFile.out")
	if err != nil {
		fmt.Println("Create traceFile:", err)
		return
	}

	defer file.Close()

	err = trace.Start(file)
	if err != nil {
		fmt.Println("Trace Start:", err)
		return
	}
	defer trace.Stop()

	var memStats runtime.MemStats

	fmt.Println("First:")
	getStats(memStats)

	for i := 0; i < 3; i++ {
		s := make([]byte, 5000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	fmt.Println("Second:")
	getStats(memStats)

	for i := 0; i < 5; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Println("Last:")
	getStats(memStats)
}

// go tool trace /tmp/traceFile.out

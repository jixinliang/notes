package main

import (
	"fmt"
	"runtime"
	"time"
)

func getMemStates(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc: ", mem.Alloc)
	fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("mem.NumGC: ", mem.NumGC)
	fmt.Println("--------------")
}


func main() {
	var mem runtime.MemStats
	getMemStates(mem)

	for i := 0; i < 10; i++ {
		_ = make([]byte, 500000)
	}

	getMemStates(mem)

	for i := 0; i < 10; i++ {
		_ = make([]byte, 1000000)
		time.Sleep(1 * time.Second)
	}

	getMemStates(mem)
}

/*
The Go compiler

go tool compile -h
go tool compile main.go
go tool compile -pack main.go

ar t main.a

go tool compile -S main.go


GODEBUG='gctrace=1' go run main.go
godebug="gctrace=1" go run main.go

4->4->0 MB
The first number is the heap size when the garbage collector is about to run.
The second value is the heap size when the garbage collector ends its operation.
The last value is the size of the live heap.
*/

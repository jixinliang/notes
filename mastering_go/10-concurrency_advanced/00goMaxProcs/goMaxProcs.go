package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GoMaxProcs:", runtime.GOMAXPROCS(0))
	fmt.Println("NumCPU:", runtime.NumCPU())
}

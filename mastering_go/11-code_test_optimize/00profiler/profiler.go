package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func fib1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(fib1(n-1)) + int64(fib1(n-2))
}

func fib2(n int) int {
	fn := make(map[int]int)

	for i := 0; i <= n; i++ {
		var tmp int
		if i <= 2 {
			tmp = 1
		} else {
			tmp = fn[i-1] + fn[i-2]
		}
		fn[i] = tmp
	}
	time.Sleep(time.Millisecond * 50)
	return fn[n]
}

func n1(n int) bool {
	k := math.Floor(float64(n))
	for i := 2; i < int(k); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func n2(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	cpuFile, err := os.Create("/tmp/cpuProfile.out")
	if err != nil {
		fmt.Println("Create cpuProfile:", err)
		return
	}

	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		fmt.Println("StartCpuProfile:", err)
		return
	}

	defer pprof.StopCPUProfile()

	total := 0
	for i := 2; i < 10000; i++ {
		n := n1(i)
		if n {
			total += 1
		}
	}
	fmt.Println("N1 Total prime number:", total)

	total = 0
	for i := 2; i < 10000; i++ {
		n := n2(i)
		if n {
			total += 1
		}
	}
	fmt.Println("N2 Total prime number:", total)

	for i := 1; i < 20; i++ {
		n := fib1(i)
		fmt.Print(n, " ")
	}

	fmt.Println()
	for i := 1; i < 20; i++ {
		n := fib2(i)
		fmt.Print(n, " ")
	}
	fmt.Println()

	runtime.GC()

	memFile, err := os.Create("/tmp/memProfile.out")
	if err != nil {
		fmt.Println("Create memProfile:", err)
		return
	}
	defer memFile.Close()

	for i := 0; i < 10; i++ {
		s := make([]byte, 500000)
		if s == nil {
			fmt.Println("Operation failed")
		}
		time.Sleep(time.Millisecond*50)
	}

	err = pprof.WriteHeapProfile(memFile)
	if err != nil {
		fmt.Println("WriteHeapProfile:",err)
		return
	}
}

// go tool pprof copProfile.out
// help
// top
// top10 --cum
// list main.n1

// sudo apt install graphviz
// go tool pprof -http=localhost:8080 cpuProfile.out

// dot -T png graph.dot -o graph.png
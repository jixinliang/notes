package main

import (
	"fmt"
	"os"
	"testing"
)

var Err error

func benchmarkCreate(b *testing.B, bufSize, fileSie int) {
	var err error
	filename := "/tmp/random.out"

	for i := 0; i < b.N; i++ {
		err = createFile(filename, bufSize, fileSie)
	}

	Err = err

	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Remove:",err)
	}
}

func Benchmark1Create(b *testing.B)  {
	benchmarkCreate(b,1,1000000)
}

func Benchmark4Create(b *testing.B)  {
	benchmarkCreate(b,4,1000000)
}

func Benchmark8Create(b *testing.B)  {
	benchmarkCreate(b,8,1000000)
}

func Benchmark100Create(b *testing.B)  {
	benchmarkCreate(b,100,1000000)
}

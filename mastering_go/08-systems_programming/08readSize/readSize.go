package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	// use make() to get the buffer
	buf := make([]byte, size)

	n, err := f.Read(buf)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println("read byte failed", err)
		return nil
	}

	return buf[0:n]
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: go run %s <buffer size> <filename>\n", filepath.Base(args[0]))
		return
	}

	bufSize, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	filename := args[2]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file failed", err)
		return
	}
	defer file.Close()

	for {
		readData := readSize(file, bufSize)
		if readData != nil {
			break
		}
		fmt.Println(string(readData))
	}
}

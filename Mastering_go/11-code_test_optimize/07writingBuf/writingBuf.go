package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

var BufferSize int
var FileSize int

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func setBuffer(buf *[]byte, size int) {
	*buf = make([]byte, size)

	if size == 0 {
		return
	}

	for i := 0; i < size; i++ {
		bytes := byte(random(0, 100))
		if len(*buf) > size {
			return
		}
		*buf = append(*buf, bytes)
	}
}

func createFile(filename string, bufSize, fileSize int) error {
	_, err := os.Stat(filename)
	if err == nil {
		return fmt.Errorf("File %s already exists.", filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes := make([]byte, 0)

	for {
		setBuffer(&bytes, bufSize)
		bytes = bytes[:bufSize]
		if _, err := file.Write(bytes); err != nil {
			return err
		}
		if fileSize < 0 {
			break
		}
		fileSize -= len(bytes)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: go run %s <BufferSize> <FileSize>\n", filepath.Base(args[0]))
		return
	}

	filename := "randomFile.txt"
	BufferSize, _ := strconv.Atoi(args[1])
	FileSize, _ := strconv.Atoi(args[2])

	err := createFile(filename, BufferSize, FileSize)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Remove:",err)
	}
}

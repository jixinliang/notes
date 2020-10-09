package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func charBychar(fileName string) error {
	var err error
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		// always read by line first
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Reading file failed", err)
			return err
		}

		// for range the string get bytes
		// transform bytes to string
		for _, c := range readString {
			fmt.Println(string(c))
		}
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("Usage: go run %s <file1> [<file2> ...]\n", filepath.Base(os.Args[0]))
		return
	}

	for _, file := range args {
		err := charBychar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func lineByline(fileName string) error {
	var err error

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Reading file failed", err)
			break
		}
		fmt.Print(readString)
	}
	return nil
}

func main() {
	flag.Parse() // before the args
	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("Usage: go run %s <file1> [<file2> ...]\n", filepath.Base(os.Args[0]))
		return
	}

	for _, file := range args {
		err := lineByline(file)
		if err != nil {
			fmt.Println("Open file failed", err)
		}
	}

}

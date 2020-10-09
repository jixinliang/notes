package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run %s column <file1> [<file2> [...<fileN]]\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	column, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Column value is not an integer:", column)
		return
	}

	if column < 0 {
		fmt.Println("Invalid column number, NUST be a positive number!")
		os.Exit(2)
	}

	for _, fileName := range args[2:] {
		fmt.Println(fileName)

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Open file failed: %s\v", err)
			continue
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Read file failed: %v\n", err)
			}

			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}

}

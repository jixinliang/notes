package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <file>\n", filepath.Base(args[0]))
		return
	}

	filename := args[1]

	fileInfo, _ := os.Stat(filename)
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size())
	fmt.Println("Mode:", fileInfo.Mode())
	fmt.Println("ModTime:", fileInfo.ModTime())
}

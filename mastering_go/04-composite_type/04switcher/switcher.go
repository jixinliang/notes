package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run %s number\n", filepath.Base(args[0]))
		os.Exit(1)
	}
	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("This value is not a number:", args[1])
	} else {
		switch {
		case num < 0:
			fmt.Println("Less than zero!")
		case num > 0:
			fmt.Println("Bigger than zero!")
		default:
			fmt.Println("Zero")
		}
	}

}

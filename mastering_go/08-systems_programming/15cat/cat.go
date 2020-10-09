package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getFile(filename string) error {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// scanning it and print its text
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	filename := ""
	args := os.Args
	// just from output to input like pipe
	if len(args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(args); i++ {
		filename = args[i]
		err := getFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}

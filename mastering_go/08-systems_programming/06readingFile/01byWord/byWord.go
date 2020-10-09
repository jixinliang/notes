package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func wordByword(fileName string) error {
	var err error
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// also read line first like line by line
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read file failed", err)
			return err
		}

		// Using regular expression split whitespace to get the word
		regu := regexp.MustCompile("[^\\s]+")
		words := regu.FindAllString(readString, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
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
		err := wordByword(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

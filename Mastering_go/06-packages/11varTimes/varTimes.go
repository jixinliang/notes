package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

var keyWord = "var"
var count = 0

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	for _, file := range args[1:] {
		fmt.Println("Processing:", file)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("read file failed", err)
			return
		}

		fileSet := token.NewFileSet()
		files := fileSet.AddFile(file, fileSet.Base(), len(content))

		var myScanner scanner.Scanner
		myScanner.Init(files, content, nil, scanner.ScanComments)

		localCount := 0
		for {
			_, tok, lit := myScanner.Scan()
			if tok == token.EOF {
				break
			}

			if lit == keyWord {
				count++
				localCount++
			}
		}
		fmt.Printf("Found %s %d times\n", keyWord, localCount)
	}
	fmt.Printf("Found %s total: %d times\n", keyWord, count)
}

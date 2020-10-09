package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	for _, file := range args[1:] {
		fmt.Println("Processing:", file)

		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Read file failed", err)
			return
		}

		fileSet := token.NewFileSet()
		files := fileSet.AddFile(file, fileSet.Base(), len(f))

		var myScanner scanner.Scanner
		myScanner.Init(files, f, nil, scanner.ScanComments)

		for {
			pos, tok, lit := myScanner.Scan()
			if tok == token.EOF {
				break
			}
			fmt.Printf("%s %s %q\n", fileSet.Position(pos), tok, lit)
		}
	}

}

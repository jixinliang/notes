package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type visitor int

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	fmt.Printf("%s %T\n", strings.Repeat("\t", int(v)), node)
	return v + 1
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}

	for _, file := range args[1:] {
		fmt.Println("Processing:", file)

		fileSet := token.NewFileSet()
		var v visitor

		f, err := parser.ParseFile(fileSet, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println("Parse failed:", err)
			return
		}

		ast.Walk(v, f)
	}
}

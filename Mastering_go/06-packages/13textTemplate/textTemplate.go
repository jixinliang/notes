package main

import (
	"fmt"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Need a template file!")
		return
	}

	fileName := args[1]

	data := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var entries []Entry

	for _, i := range data {
		if len(i) == 2 {
			tmp := Entry{Number: i[0], Square: i[1]}
			entries = append(entries, tmp)
		}
	}

	t := template.Must(template.ParseGlob(fileName))
	t.Execute(os.Stdout, entries)
}

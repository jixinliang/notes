package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <example.csv>\n", filepath.Base(args[0]))
		return
	}

	filename := args[1]
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Println("file does not exist!")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file failed!")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Read file failed!")
		return
	}

	for _, rec := range records {
		fmt.Println("Record:", rec)
	}
}

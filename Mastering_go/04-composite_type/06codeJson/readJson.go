package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Record struct {
	Name    string
	SurName string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJson(fileName string, key interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(key)
	if err != nil {
		return err
	}
	return nil
}

func loadDemo() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run %s jsonFile\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	fileName := args[1]

	var record Record
	err := loadFromJson(fileName, &record)
	if err != nil {
		fmt.Println("Load json Faild:", err)
	}
	fmt.Println(record)
}

func saveToJson(fileName *os.File, key interface{}) {
	encoder := json.NewEncoder(fileName)
	err := encoder.Encode(key)
	if err != nil {
		fmt.Println("Encoder json faild:", err)
		return
	}
}

func main() {
	record := Record{
		Name:    "Jack",
		SurName: "Ji",
		Tel: []Telephone{
			{Mobile: true, Number: "1234-567"},
			{Mobile: true, Number: "1234-789"},
			{Mobile: true, Number: "1234-abc"},
		},
	}
	saveToJson(os.Stdout, record)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run %s jsonFile\n", filepath.Base(args[0]))
		return
	}

	fileName := args[1]
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read file failed:", err)
		return
	}

	var parsedData map[string]interface{}

	json.Unmarshal(fileBytes, &parsedData)
	for k, v := range parsedData {
		fmt.Println("key:", k, "val:", v)
	}
}

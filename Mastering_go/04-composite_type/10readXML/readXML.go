package main

import (
	"encoding/xml"
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

func loadFromXML(fileName string, key interface{}) error {
	// 1.open file
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	// 2. delay close the file
	defer file.Close()

	// 3. create xmlDecoer
	decoder := xml.NewDecoder(file)

	// 4. decode the file to that struct
	err = decoder.Decode(key)
	if err != nil {
		return err
	}

	// 5. if error equal nil
	return nil

}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run %s xmlFile\n", filepath.Base(args[0]))
		return
	}

	fileName := args[1]

	var record Record

	err := loadFromXML(fileName, &record)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(record)
}

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type Record struct {
	Name    string      `json:"name"`
	SurName string      `json:"sur_name"`
	Tel     []Telephone `json:"tel"`
}

type Telephone struct {
	Mobile bool   `json:"mobile"`
	Number string `json:"number"`
}

// Parsing file to go json data struct
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

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run %s jsonFile\n", filepath.Base(args[0]))
		return
	}

	fileName := args[1]

	var record Record

	err := loadFromJson(fileName, &record)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(record)

	xmlBytes, _ := xml.MarshalIndent(record, "", "    ")

	// Parsing json to xml
	xmlBytes = []byte(xml.Header + string(xmlBytes))
	fmt.Println("xml:", string(xmlBytes))

	// Parsing xml to json
	var jsonRecord Record
	err = xml.Unmarshal(xmlBytes, &jsonRecord)
	if err != nil {
		fmt.Println("Unmarshal xml failed", err)
		return
	}

	resBytes, err := json.Marshal(jsonRecord)
	if err != nil {
		fmt.Println("Marshal record faild", err)
		return
	}

	json.Unmarshal(resBytes, &record)
	fmt.Println("Json:", record)

}

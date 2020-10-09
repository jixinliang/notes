package main

import (
	"encoding/json"
	"fmt"
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

	jsonBytes, err := json.Marshal(&record)
	if err != nil {
		fmt.Println("Marshal json faild:", err)
		return
	}
	//fmt.Println(string(jsonBytes))

	var unRecord Record
	err = json.Unmarshal(jsonBytes, &unRecord)
	if err != nil {
		fmt.Println("Unmarshal json faild:", err)
		return
	}
	fmt.Println(unRecord)
}

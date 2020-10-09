package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Address struct {
	City    string `xml:"city"`
	Country string `xml:"country"`
}

type Employee struct {
	XmlName   xml.Name `xml:"xml_name"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"first_name"`
	LastName  string   `xml:"last_name"`
	Initials  string   `xml:"initials"`
	Height    float32  `xml:"height"`
	Address
	Comment string `xml:"comment"`
}

func main() {
	e := &Employee{ID: 1, FirstName: "Jack", LastName: "Ji", Initials: "MIT", Height: 182}
	e.Address = Address{City: "AnHi", Country: "China"}
	e.Comment = "Technical Writer"

	xmlBytes, err := xml.MarshalIndent(e, "  ", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	xmlBytes = []byte(xml.Header + string(xmlBytes))
	os.Stdout.Write(xmlBytes)
	os.Stdout.Write([]byte("\n"))
}

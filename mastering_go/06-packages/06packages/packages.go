package main

import (
	"06packages/package1"
	"fmt"
)

func main() {
	fmt.Println("Using packageA:")
	package1.Msg()
	fmt.Println("myConst:", package1.MyConst)
	println("Hello World")
	print("Hello there!\n")
}

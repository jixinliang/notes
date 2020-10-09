package main

import (
	"07init/func1"
	"07init/func2"
	"fmt"
)

func init() {
	fmt.Println("package main init()")
}

func main() {
	fmt.Println("calling function main()")
	func1.FuncA()
	func2.FuncB()
}

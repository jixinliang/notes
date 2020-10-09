package main

import "fmt"

func saiHi() {
	fmt.Println("Hi")
}

func main() {
	saiHi()
}

// GOSSAFUNC=man go build -gcflags="-S" toMachineCode.go
// Static Single Assignment (SSA)
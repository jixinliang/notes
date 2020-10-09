package main

import "fmt"

const C1 = 123
const C2 string = "c2"

const (
	Zero int = iota
	One
	Two
	Three
	Four
)

const(
	p0 int = 1 << iota
	_ // skip unwanted values.
	p2
	_
	p4
)

func main() {
	fmt.Printf("type of c1: %T\n", C1)
	fmt.Printf("type of c2: %T\n", C2)

	var v1 float32 = C1 + 10
	fmt.Println("v1:", v1)

	fmt.Println("one:",One)
	fmt.Println("four:",Four)
	fmt.Println(p0, p2, p4)
}

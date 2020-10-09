package main

import "fmt"

type square struct {
	X float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func assertType(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("This is a square!")
	case circle:
		fmt.Printf("%v is a circle!\n", v)
	case rectangle:
		fmt.Println("This is a rectangle!")
	default:
		fmt.Printf("Unknow type %T!\n", v)
	}
}

func main() {
	c := circle{10}
	assertType(c)

	s := square{4}
	assertType(s)

	r := rectangle{3,5}
	assertType(r)

	assertType(0)
}

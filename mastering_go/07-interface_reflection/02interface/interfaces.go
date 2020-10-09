package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	X float64
}

type Circle struct {
	R float64
}

func (s Square) Area() float64 {
	return s.X * s.X
}

func (s Square) Perimeter() float64 {
	return s.X * 4
}

func (c Circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func Calc(x Shape) {
	_, ok := x.(Circle)
	if ok {
		fmt.Println("Is a Circle")
	}

	s, ok := x.(Square)
	if ok {
		fmt.Println("Is a Square", s)
	}
	fmt.Println("Area:", x.Area())
	fmt.Println("Perimeter:", x.Perimeter())
}

func main() {
	s := Square{10}
	fmt.Println("Perimeter:", s.Perimeter())
	Calc(s)

	c := Circle{5}
	Calc(c)
}

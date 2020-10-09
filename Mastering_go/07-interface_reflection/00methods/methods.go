package main

import "fmt"

type Ints struct {
	X int64
	Y int64
}

func regularFunc(a, b Ints) Ints {
	tmp := Ints{a.X + b.X, a.Y + b.Y}
	return tmp
}

func (i Ints) method(j Ints) Ints {
	tmp := Ints{i.X + j.X, i.Y + j.Y}
	return tmp
}

// The method() function is equivalent to the regularFunc() function

func main() {
	i := Ints{1, 2}
	j := Ints{-5, -2}
	fmt.Println("RegularFunction:", regularFunc(i, j))
	fmt.Println("Method:", i.method(j))
}

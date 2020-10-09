package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	a := [...]int{4, 5, 6}

	fmt.Printf("type of a: %T\n", a)
	res := a[:]
	fmt.Println("array transform to slice:", res)
	fmt.Printf("type of res: %T\n", res)
	s1 := append(s, res...)
	fmt.Println("s1+res:", s1)

	s = append(s, res...)
	fmt.Println("s+res:", s)

	s = append(s, s...)
	fmt.Println("s+s:", s)

}

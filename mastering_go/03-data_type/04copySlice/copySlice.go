package main

import "fmt"

func long2Short() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6, 7, 8}
	fmt.Println("before copy:")
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	copy(a1, a2)
	fmt.Println("after copied:")
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
}

func long2Short1() {
	a1 := [3]int{1, 2, 3}
	a2 := []int{4, 5, 6, 7, 8}
	fmt.Println("before copy:")
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	copy(a1[0:], a2)
	fmt.Println("after copied:")
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
}

func short2Long() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6, 7, 8}
	fmt.Println("before copy:")
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	copy(a2, a1)
	fmt.Println("after copied:")
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
}

func twoDSlice() {
	s := make([][]int, 3)
	fmt.Println(s)
	for _, i := range s {
		for k, v := range i {
			fmt.Println("k:", k, "v:", v)
		}
		fmt.Println()
	}
}

func main() {
	twoDSlice()
}

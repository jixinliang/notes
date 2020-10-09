package main

import "fmt"

func getSliceVal(s []int) {
	for _, val := range s {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func main() {
	s1 := []int{1, 2, 3}
	fmt.Print("s1: ")
	getSliceVal(s1)
	fmt.Printf("cap: %d, len: %d\n", cap(s1), len(s1))

	fmt.Println("--------------------------------------------")
	s1 = append(s1, -1)
	fmt.Print("s1: ")
	getSliceVal(s1)
	fmt.Printf("cap: %d, len: %d\n", cap(s1), len(s1))

	fmt.Println("--------------------------------------------")
	s1 = append(s1, -1)
	s1 = append(s1, -2)
	s1 = append(s1, -3)
	fmt.Print("s1: ")
	getSliceVal(s1)
	fmt.Printf("cap: %d, len: %d\n", cap(s1), len(s1))

}

//  slices are dynamic in size, if a slice runs out of room,
// Go automatically doubles its current length to make room for more elements.
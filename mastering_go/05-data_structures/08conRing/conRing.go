package main

import (
	"container/ring"
	"fmt"
)

func main() {
	size := 10

	myRing := ring.New(size+1)
	fmt.Println("Empty ring:", *myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	myRing.Value = 2
	//fmt.Println("Result1 ring:", *myRing)

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum += t
	})
	fmt.Println("Sum:", sum)

	// imply the value is a ring
	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
}

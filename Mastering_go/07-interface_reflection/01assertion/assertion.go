package main

import "fmt"

func main() {
	var myInt interface{} = 123
	i, ok := myInt.(int)
	if ok {
		fmt.Println("Assert success, it's an int:", i)
	}

	f, ok := myInt.(float64)
	if ok {
		fmt.Println(f)
	} else {
		fmt.Println("Assert failed it not float64 type")
	}

	i = myInt.(int)
	fmt.Println("No checking:", i)

	b := myInt.(bool)
	fmt.Println(b)
}

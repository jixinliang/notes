package main

import (
	"fmt"
	"os"
)

func f1() {
	fmt.Println("Inside f1")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover inside f1")
		}
	}()

	fmt.Println("to call f2")
	f2()
	fmt.Println("f2() exited")
	fmt.Println("f1() exited")
}

func f2() {
	fmt.Println("Inside f2")
	panic("Panic in f2()")
}

func caller() {
	f1()
	fmt.Println("function main()")
}

func notEnough() {
	if len(os.Args) == 1 {
		panic("Not enough arguments")
	}
	fmt.Println("Thanks for the arguments")
}

func main() {
	notEnough()
}

/*
panic() is a built-in Go function that terminates the current flow of a Go
program and starts panicking. On the other hand, the recover() function, which is also a
built-in Go function, allows you to take back control of a goroutine that just panicked
using panic().
*/

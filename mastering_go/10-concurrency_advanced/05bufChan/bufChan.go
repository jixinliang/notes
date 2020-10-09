package main

import "fmt"

func main() {
	nums := make(chan int, 5)
	counter := 10

	// receive the data
	for i := 0; i < counter; i++ {
		select {
		case nums <- i: // receiver
		default:
			fmt.Println("Not enough space for:", i)
		}
	}

	// send data
	for i := 0; i < counter+5; i++ {
		select {
		case num := <-nums: // sender
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done.")
			break
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	c <- -1
	c <- 0
	c <- 1

	<-c
	<-c
	//<-c
	fmt.Println("Before close:", <-c)
	close(c)

	data := <-c
	// reading from a closed channel will get a
	// zero value of the data type
	fmt.Println("data:", data)
}

package main

import (
	"fmt"
	"time"
)

func writeToChan(c chan int, x int) {

	fmt.Println("Before:", x)
	c <- x
	close(c)
	fmt.Println("After:", x)

}

func main() {
	c := make(chan int)
	go writeToChan(c, 10)
	time.Sleep(time.Second)
}

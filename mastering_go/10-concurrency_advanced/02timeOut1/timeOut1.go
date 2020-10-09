package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "C1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("Timeout c1") // sleep time less then massage
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "C2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout c2")
	}
}

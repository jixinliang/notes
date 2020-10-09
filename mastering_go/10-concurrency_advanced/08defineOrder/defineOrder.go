package main

import (
	"fmt"
	"time"
)

func f1(a, b chan struct{}) {
	<-a
	fmt.Println("f1()!")
	time.Sleep(time.Second)
	close(b)
}

func f2(a, b chan struct{}) {
	<-a
	fmt.Println("f2()!")
	close(b)
}

func f3(a chan struct{}) {
	<-a
	fmt.Println("f3()!")
}

func main() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})

	go f3(c3)
	go f1(c1, c2)
	go f3(c3)
	go f2(c2, c3)
	go f3(c3)
	close(c1)
	time.Sleep(time.Second * 3)
}

/*
The main advantage of a struct{} signal channel is that no data can be sent to it,
which can save you from bugs and misconceptions
*/
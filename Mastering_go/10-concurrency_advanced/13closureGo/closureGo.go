package main

import (
	"fmt"
	"time"
)

func unCorrect()  {
	for i := 0; i <= 20; i++ {
		go func() {
			fmt.Print(i," ")
		}()
	}

	time.Sleep(time.Second)
	fmt.Println()
}

func correct()  {
	for i := 0; i <= 20; i++ {
		i := i
		go func() {
			fmt.Print(i," ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}




func main() {



}

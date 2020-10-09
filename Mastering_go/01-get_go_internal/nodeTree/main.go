package main

import "fmt"

func main() {
	fmt.Println("Hello There!")
}

// go tool compile -W main.go
// If your program had more functions, you would have got more output,
// go tool compile -W main.go|grep before
// go tool compile -W main.go|grep after

// go build -x main.go

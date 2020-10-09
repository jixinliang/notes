package main

import "fmt"

func add(x, y uint16) uint16 {
	var i uint16
	for i = 0; i < x; i++ {
		y++
	}
	return y
}

func main() {
	fmt.Println(add(0,0))
}

// go test -v quick*
// go test -v quick* -count=1
// go help testflag
package main

import (
	"flag"
	"fmt"
)

func main() {
	k := flag.Bool("k", true, "k flag")
	o := flag.Int("O", 1, "O")
	flag.Parse()

	valK := *k
	valO := *o
	valO++
	fmt.Println("Value k:", valK)
	fmt.Println("Value o:", valO)

}

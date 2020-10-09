package main

import "fmt"

func main() {
	arr := [12][4][7][10]float64{}
	x := len(arr)
	y := len(arr[0])
	z := len(arr[0][0])
	w := len(arr[0][0][0])
	fmt.Println("X:", x, "Y:", y, "Z:", z, "W:", w)
}

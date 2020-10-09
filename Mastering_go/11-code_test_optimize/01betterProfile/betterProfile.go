package main

import (
	"fmt"
	"github.com/pkg/profile"
)

var Variable int

func n1(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func multi(a, b int) int {
	if a == 1 {
		return b
	}
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		return -multi(-a, b)
	}
	return b + multi(a-1, b)
}

func main() {
	defer profile.Start(profile.ProfilePath("/tmp")).Stop()
	defer profile.Start(profile.MemProfile).Stop()

	total := 0

	for i := 2; i < 1000; i++ {
		n := n1(i)
		if n {
			total += 1
		}
	}
	fmt.Println("Total prime number:", total)

	total = 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 40; j++ {
			res := multi(i, j)
			Variable = res
			total++
		}
	}
	fmt.Println("Total mulri number:", total)
}

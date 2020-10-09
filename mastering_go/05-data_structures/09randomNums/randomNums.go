package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	min := 0
	max := 100
	total := 100
	seed := time.Now().Unix()
	args := os.Args

	switch len(args) {
	case 2:
		fmt.Printf("Usage: ./%s min max total seed\n", filepath.Base(args[0]))
		min, _ = strconv.Atoi(args[1])
		max = min + 100
	case 3:
		fmt.Printf("Usage: ./%s min max total seed\n", filepath.Base(args[0]))
		min, _ = strconv.Atoi(args[1])
		max, _ = strconv.Atoi(args[2])
	case 4:
		fmt.Printf("Usage: ./%s min max total seed\n", filepath.Base(args[0]))
		min, _ = strconv.Atoi(args[1])
		max, _ = strconv.Atoi(args[2])
		total, _ = strconv.Atoi(args[3])
	case 5:
		fmt.Printf("Usage: ./%s min max total seed\n", filepath.Base(args[0]))
		min, _ = strconv.Atoi(args[1])
		max, _ = strconv.Atoi(args[2])
		total, _ = strconv.Atoi(args[3])
		seed, _ = strconv.ParseInt(args[4], 10, 64)
	default:
		fmt.Println("Using default values.")
	}

	rand.Seed(seed)
	for i := 0; i < total; i++ {
		myRand := random(min, max)
		fmt.Print(myRand, " ")
	}
	fmt.Println()
}

/*
Go uses the math/rand package for generating pseudo-random numbers. It needs a seed to
start producing the numbers. The seed is used for initializing the entire process, and it is
extremely important because if you always start with the same seed, you will always get
the same sequence of pseudo-random numbers. This means that everybody can regenerate
that sequence, and that particular sequence will not be random after all
*/

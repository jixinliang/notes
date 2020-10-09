package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func multiMatrices(m1, m2 [][]int) ([][]int, error) {
	if len(m1[0]) != len(m2) {
		return nil, errors.New("cannot multiply the given matrices")
	}
	res := make([][]int, len(m1))
	for i := 0; i < len(m1); i++ {
		res[i] = make([]int, len(m2[0]))
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res, nil
}

func createMatrix(row, col int) [][]int {
	res := make([][]int, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res[i] = append(res[i], random(-5, i*j))
		}
	}
	return res
}

func main() {
	args := os.Args
	rand.Seed(time.Now().Unix())
	if len(args) != 5 {
		fmt.Println("Wrong number of arguments!")
		return
	}

	var row, col int

	// for m1
	row, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Need an integer:", args[1])
		return
	}

	col, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Need an integer:", args[2])
		return
	}

	if row <= 0 || col <= 0 {
		fmt.Println("Need positive matrix diamensions!")
		return
	}

	fmt.Printf("m1 is %d x %d matrix\n", row, col)
	m1 := createMatrix(row, col)

	// for m2
	row, err = strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Need an integer:", args[3])
		return
	}

	col, err = strconv.Atoi(args[4])
	if err != nil {
		fmt.Println("Need an integer:", args[4])
		return
	}

	if row <= 0 || col <= 0 {
		fmt.Println("Need positive matrix diamonsons!")
		return
	}
	fmt.Printf("m2 is %d x %d matrix\n", row, col)
	m2 := createMatrix(row, col)
	fmt.Println("M1:", m1)
	fmt.Println("M2:", m2)

	// Multiply
	r1, err := multiMatrices(m1, m2)
	if err != nil {
		fmt.Println("Mutiply failed,", err)
		return
	}
	fmt.Println("R1:", r1)
}

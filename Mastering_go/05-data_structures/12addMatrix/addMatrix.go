package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// replace with its opposite integer
func nagativeMatrix(s [][]int) [][]int {
	for i, x := range s {
		for j, _ := range x {
			s[i][j] = -s[i][j]
		}
	}
	return s
}

// add two matrixces into one matrix
func addMatrixces(m1, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i, x := range m1 {
		for j, _ := range x {
			res[i] = append(res[i], m1[i][j]+m2[i][j])
		}
	}
	return res
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Wrong number of arguments")
		return
	}

	var row, col int
	row, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Require a integer:", args[1])
		return
	}

	col, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Require a integer:", args[2])
		return
	}

	fmt.Printf("Using %d x %d arrays\n", row, col)

	if col <= 0 || row <= 0 {
		fmt.Println("Require positive matrix dimensions.")
		return
	}

	m1 := make([][]int, row)
	m2 := make([][]int, row)

	rand.Seed(time.Now().Unix())

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m1[i] = append(m1[i], random(-1, i*j+rand.Intn(10)))
			m2[i] = append(m2[i], random(-1, i*j+rand.Intn(10)))
		}
	}

	fmt.Println("M1:", m1)
	fmt.Println("M2:", m2)

	// add
	r1 := addMatrixces(m1, m2)
	// subtract
	r2 := addMatrixces(m1, nagativeMatrix(m2))
	fmt.Println("R1 for add:", r1)
	fmt.Println("R2 for substract:", r2)

}

/*
There are some rules that can tell you whether you can perform a calculation between two
matrices or not. The rules are the following:
In order to add or subtract two matrices, they should have exactly the same
dimensions.
In order to multiply matrix A with matrix B, the number of columns of matrix A
should be equal to the number of rows of matrix B. Otherwise, the multiplication
of matrices A and B is impossible.
In order to divide matrix A with matrix B, two conditions must be met. Firstly,
you will need to be able to calculate the inverse of matrix B and secondly, you
should be able to multiply matrix A with the inverse of matrix B according to the
previous rule. Only square matrices can have an inverse
*/

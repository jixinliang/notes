package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func importFile(fileName string) ([][]int, error) {
	var err error
	var S = make([][]int, 0)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		fields := strings.Fields(readString)
		tmp := make([]int, 0)

		for _, v := range fields {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, n)
		}
		if len(tmp) != 0 {
			S = append(S, tmp)
		}

		if len(tmp) != len(S[0]) {
			return nil, errors.New("wrong number of elements")
		}
	}
	return S, nil
}

func validPuzzle(S [][]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; i <= 2; j++ {
			i1 := i * 3
			j1 := j * 3
			s := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					i2 := i1 + k
					j2 := j1 + m
					val := S[i2][j2]
					if val > 0 && val < 10 {
						if s[val-1] == 1 {
							fmt.Println("Appeared 2 times:", val)
							return false
						} else {
							s[val-1] = 1
						}
					} else {
						fmt.Println("Invalid value:", val)
						return false
					}
				}
			}
		}
	}

	// testing columns
	for i := 0; i <= 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum += S[i][j]
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}

	// testing rows
	for i := 0; i < 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum += S[j][i]
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}
	return true
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: go run %s testFile size\n", filepath.Base(args[0]))
		return
	}

	fileName := args[1]
	ss, err := importFile(fileName)
	if err != nil {
		fmt.Println("Import file failed", err)
		return
	}

	if validPuzzle(ss) {
		fmt.Println("Correct Sudoku puzzle.")
	} else {
		fmt.Println("Incorrect Sudoku puzzle!")
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func min(args []float64) float64 {
	return args[0]
}

func max(args []float64) float64 {
	return args[len(args)-1]
}

func mean(args []float64) float64 {
	sum := float64(0)
	for _, val := range args {
		sum += val
	}
	return sum / float64(len(args))
}

func median(args []float64) float64 {
	length := len(args)
	// Odd
	if length%2 == 1 {
		return args[(length-1)/2]
	} else {
		// Even
		return (args[length/2] + args[(length/2)-1]) / 2
	}
	return 0
}

func variance(args []float64) float64 {
	meanVal := mean(args)
	sum := float64(0)
	for _, val := range args {
		// 每个样本值与全体样本值的平均数之差的平方平均数
		sum += (val - meanVal) * (val - meanVal)
	}
	return sum / float64(len(args))
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("Usage: go run %s <file.txt>\n", filepath.Base(os.Args[0]))
		return
	}

	filename := args[0]
	data := make([]float64, 0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error at Open()->", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error at Read file->", err)
			return
		}
		readString = strings.TrimRight(readString, "\r\n")
		val, err := strconv.ParseFloat(readString, 64)
		if err != nil {
			fmt.Println("Error at ParseFloat()->", err)
			return
		}
		data = append(data, val)
	}
	sort.Float64s(data)

	fmt.Println("Min:", min(data))
	fmt.Println("Max:", max(data))
	fmt.Println("Mean:", mean(data))
	fmt.Println("Median:", median(data))
	fmt.Println("Variance:", variance(data))
	fmt.Println("Standard Deviation:", math.Sqrt(variance(data)))
}

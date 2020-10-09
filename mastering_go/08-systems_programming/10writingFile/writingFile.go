package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var msg = []byte("Hello There!")

func wf1() {
	// way1
	file1, err := os.Create("file1.txt")
	if err != nil {
		fmt.Println("Create file1 failed!")
		return
	}
	defer file1.Close()

	// way 1 key
	fmt.Fprint(file1, string(msg))
}

func wf2() {
	// way2
	file2, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println("Create file2 failed!")
		return
	}
	defer file2.Close()

	n, err := file2.WriteString(string(msg))
	if err != nil {
		fmt.Println("Write string to file2 failed!")
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
}

func wf3() {
	// way3
	file3, err := os.Create("file3.txt")
	if err != nil {
		fmt.Println("Create file3 faild!")
		return
	}

	writer := bufio.NewWriter(file3)
	n, err := writer.WriteString(string(msg))
	if err != nil {
		fmt.Println("Write string to file3 failed!")
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
	writer.Flush()
}

func wf4() {
	// way 4
	file4 := "file4.txt"
	err := ioutil.WriteFile(file4, msg, 0644)
	if err != nil {
		fmt.Println("Write to file4 failed!")
		return
	}
}

func wf5() {
	// way 5
	file5, err := os.Create("file5.txt")
	if err != nil {
		fmt.Println("Create file5 faild!")
		return
	}

	n, err := io.WriteString(file5, string(msg))
	if err != nil {
		fmt.Println("Write string to file5 failed!")
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
}

func main() {

}

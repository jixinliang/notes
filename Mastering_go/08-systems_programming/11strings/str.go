package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	msg := "Hello World!"
	reader := strings.NewReader(msg)
	fmt.Println("length of msg:", reader.Len())

	b := make([]byte, 1)
	for {
		n, err := reader.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read bytes failed", err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}

	errMsg := "This is an error message\n"
	newReader := strings.NewReader(errMsg)
	fmt.Println("length of errMsg:", newReader.Len())

	n, err := newReader.WriteTo(os.Stderr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d bytes to StdErr\n", n)
}

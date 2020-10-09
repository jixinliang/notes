package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer

	msg := []byte("Byte!\n")
	buffer.Write(msg)
	fmt.Fprintf(&buffer, "a string!\n")

	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stdout)

	buffer.Reset()

	msg1 := []byte("Hi")
	buffer.Write(msg1)

	reader := bytes.NewReader([]byte(buffer.String()))
	fmt.Println("buffer string:", buffer.String())

	for {
		b := make([]byte, 3)
		n, err := reader.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read byte failed:", err)
			return
		}
		fmt.Printf("Read %s bytes: %d", b, n)
	}

}

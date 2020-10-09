package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Domain Name>\n", filepath.Base(args[0]))
		return
	}

	mxes, err := net.LookupMX(args[1])
	if err != nil {
		fmt.Println("Error in LookupMX:", err)
		return
	}

	for _, mx := range mxes {
		fmt.Println(mx.Host)
	}
}

// host -t mx golang.org

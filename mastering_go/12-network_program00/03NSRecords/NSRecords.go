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

	NSs, err := net.LookupNS(args[1])
	if err != nil {
		fmt.Println("Error in LookupNS:", err)
		return
	}
	for _, ns := range NSs {
		fmt.Println(args[1], "name server", ns.Host)
	}
}

// host -t ns bing.com

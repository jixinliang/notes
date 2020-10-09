package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Compiler:", runtime.Compiler)
	fmt.Println("GOArch:", runtime.GOARCH)
	fmt.Println("Version:", runtime.Version())
}

// GOARCH=386 go run crossCompile.go
// GOOS=windows GOARCH=386 go run crossCompile.go

package main

import "fmt"

func main() {
	fmt.Println("Create WebAssembly code from go")
}

// GOOS=js GOARCH=wasm go build -o main.wasm toWasm.go
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

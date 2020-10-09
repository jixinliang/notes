package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <URL>\n", filepath.Base(args[0]))
		return
	}

	url := args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in Get:", err)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println("Error in Copy:", err)
		return
	}

}

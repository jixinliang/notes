package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var timeout = time.Duration(time.Second)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <URL>\n", filepath.Base(args[0]))
		return
	}

	if len(args) == 3 {
		tmp, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Using default Timeout", timeout)
		} else {
			timeout = time.Duration(time.Duration(tmp) * time.Second)
			fmt.Println("Using Timeout:", timeout)
		}
	}

	urlStr := args[1]

	client := http.Client{Timeout: timeout}
	resp, err := client.Get(urlStr)
	if err != nil {
		fmt.Println("Error at Get:", err)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println("Error at Copy:", err)
		return
	}
}

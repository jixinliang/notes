package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <URL>\n", filepath.Base(args[0]))
		return
	}

	url := args[1]

	transport := &http.Transport{TLSClientConfig: &tls.Config{}}

	client := &http.Client{Transport: transport}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error at Get:", err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error at ReadAll:", err)
		return
	}
	ctx := strings.TrimSpace(string(content))
	fmt.Println(ctx)
}

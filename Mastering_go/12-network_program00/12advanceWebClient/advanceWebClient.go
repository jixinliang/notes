package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <http://example.com>\n", filepath.Base(args[0]))
		return
	}

	urlObj, err := url.Parse(args[1])
	if err != nil {
		fmt.Println("Error at Parse:", err)
		return
	}

	request, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Println("Error at NewRequest:", err)
		return
	}

	client := &http.Client{Timeout: 15 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error at Do:", err)
		return
	}

	fmt.Println("Status:", response.Status)

	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		fmt.Println("Error at DumpResponse", err)
		return
	}
	fmt.Println("Dump:", string(dump))

	contentType := response.Header.Get("Content-Type")
	charSet := strings.SplitAfter(contentType, "charset=")
	if len(charSet) > 1 {
		fmt.Println("CharSet:", charSet[1])
	}

	if response.ContentLength != -1 {
		fmt.Println(response.ContentLength)
	} else {
		fmt.Println("ContentLength UnKnown!")
	}

	length := 0
	var buffer [1024]byte

	r := response.Body
	for {
		_, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println("Error at Read:", err)
			break
		}
		length += 1
	}

	fmt.Println("Length of data of Body:",length)
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
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

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error at NewRequest:", err)
		return
	}

	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("Got First Response Byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Println("GotConnInfo:", connInfo)
		},
		DNSDone: func(dnsDoneInfo httptrace.DNSDoneInfo) {
			fmt.Println("DNSDoneInfo:", dnsDoneInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Connect Start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Connect Done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote Header")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Requesting data from server!")

	_, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println("Error in RoundTrip")
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in Do:", err)
		return
	}

	io.Copy(os.Stdout, res.Body)
}

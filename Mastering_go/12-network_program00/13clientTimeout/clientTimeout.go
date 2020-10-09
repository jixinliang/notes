package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var duration = time.Duration(time.Second)

func timeout(network, address string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, address, duration)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(duration))
	return conn, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <URL>\n", filepath.Base(args[0]))
		return
	}

	if len(args) == 3 {
		tmp, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Using default Timeout!")
		} else {
			duration = time.Duration(time.Duration(tmp) * time.Second)
		}
	}

	url := args[1]

	trans := http.Transport{Dial: timeout}

	client := http.Client{Transport: &trans}
	resp, err := client.Get(url)
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

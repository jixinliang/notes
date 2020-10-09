package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Port>\n", filepath.Base(args[0]))
		return
	}

	addr := "localhost" + ":" + args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println("Error at ResolveTCPAddr()->", err)
		return
	}

	tcpListener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println("Error at ListenTCP()->", err)
		return
	}

	conn, err := tcpListener.Accept()
	if err != nil {
		fmt.Println("Error at Accept()->", err)
		return
	}

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error at Read()->", err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "exit" {
			fmt.Println("Exiting TCP server")
			conn.Close()
			break
		}

		fmt.Print("> ", string(buffer[0:n-1]))

		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println("Error at Write()->", err)
			return
		}
	}
}

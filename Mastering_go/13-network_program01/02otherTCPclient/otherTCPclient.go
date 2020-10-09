package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Addr:port>\n", filepath.Base(args[0]))
		return
	}

	addr := args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println("Error at ResolveTCPAddr()->", err)
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("Error at DialTCP()->", err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error at Read Stdin->", err)
			return
		}
		fmt.Fprintf(conn, text, "\n")

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error at Read conn->", err)
			return
		}
		fmt.Print("> ", msg)

		if strings.TrimSpace(string(msg)) == "exit" {
			fmt.Println("TCP client exiting...")
			conn.Close()
			break
		}
	}
}

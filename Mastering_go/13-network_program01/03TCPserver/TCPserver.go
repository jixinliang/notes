package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Port>\n", filepath.Base(args[0]))
		return
	}

	port := ":" + args[1]
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error at Listen()->", err)
		return
	}
	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Error at Accept()->", err)
		return
	}

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error at Read conn->", err)
			return
		}
		if strings.TrimSpace(string(data)) == "exit" {
			fmt.Println("Exiting TCP server")
			return
		}

		fmt.Print("> ", string(data))
		nowTime := time.Now().Format(time.RFC3339) + "\n"
		conn.Write([]byte(nowTime))
	}
}

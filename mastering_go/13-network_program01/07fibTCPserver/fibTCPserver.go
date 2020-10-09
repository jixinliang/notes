package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func fib(n int) int {
	fn := make(map[int]int)

	// note i must little equal n
	for i := 0; i <= n; i++ {
		var val int
		if i <= 2 {
			val = 1
		} else {
			val = fn[i-1] + fn[i-2]
		}
		fn[i] = val
	}
	return fn[n]
}

func connHandler(conn net.Conn) {
	for {
		readString, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error at read conn->", err)
			return
		}

		tmp := strings.TrimSpace(string(readString))

		if tmp == "exit" {
			break
		}

		fibo := "-1" + "\n"
		n, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Println("Error at Atoi()->", err)
			return
		}
		fibo = strconv.Itoa(fib(n)) + "\n"

		conn.Write([]byte(string(fibo)))
	}
	time.Sleep(time.Second * 2)
	conn.Close()
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Port>\n", filepath.Base(args[0]))
		return
	}

	port := ":" + args[1]

	addr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		fmt.Println("Error at ResolveTCPserver()->", err)
		return
	}

	tcpListener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println("Error at ListenTCP()->", err)
		return
	}

	defer tcpListener.Close()

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println("Error at Accept()->", err)
			return
		}
		go connHandler(conn)
	}
}

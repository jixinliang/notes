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

	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		fmt.Println("Error at ResolveUDPAddr()->", err)
		return
	}

	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		fmt.Println("Error at DialUDP()->", err)
		return
	}
	defer conn.Close()
	fmt.Printf("The UDP server is %s\n", conn.RemoteAddr().String())

	for {
		// read form stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error at read Stdin->", err)
			return
		}

		data := []byte(text + "\n")

		// write
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error at Write()->", err)
			return
		}

		if strings.TrimSpace(string(data)) == "exit" {
			fmt.Println("Exiting UPD server...")
			break
		}

		// read from server
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error at ReadFromUDP()->", err)
			return
		}

		fmt.Printf("Reply: %s\n", string(buffer[:n]))
	}
}

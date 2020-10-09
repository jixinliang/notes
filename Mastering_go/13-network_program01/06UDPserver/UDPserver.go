package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Port>\n", filepath.Base(args[0]))
		return
	}

	port := ":" + args[1]

	// 1. resolve port
	addr, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println("Error at ResolveUDPAddr()->", err)
		return
	}

	// listen addr
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("Error at ListenUDP()->", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		// 3. read data
		n, udpAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error at ReadFromUDP()->", err)
			return
		}

		if strings.TrimSpace(string(buffer[:n])) == "exit" {
			fmt.Println("Exiting UDP server...")
			break
		}

		// send data
		data := []byte(strconv.Itoa(random(1, 1024)))
		fmt.Println("Data:", string(data))
		_, err = conn.WriteToUDP(data, udpAddr)
		if err != nil {
			fmt.Println("Error at WriteToUDP()->", err)
			return
		}
	}

}

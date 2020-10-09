package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Println("Error at Socket()->", err)
		return
	}

	file := os.NewFile(uintptr(fd), "captureICMP")
	if file == nil {
		fmt.Println("Error at NewFile()->")
		return
	}
	defer file.Close()

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, 256)
	if err != nil {
		fmt.Println("Error at SetsockoptInt()->", err)
		return
	}

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("Error at Read()->", err)
			return
		}

		fmt.Printf("% X\n", buffer[:n])
	}
}

// ping localhost
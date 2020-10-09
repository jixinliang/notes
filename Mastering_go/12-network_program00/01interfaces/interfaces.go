package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Interfaces:", err)
		return
	}

	for _, i := range interfaces {
		fmt.Println("Name:", i.Name)
		fmt.Println("Flags:", i.Flags.String())
		fmt.Println("HardwareAddr:", i.HardwareAddr)
		fmt.Println("MTU:", i.MTU)

		fmt.Println()
	}
}

// netstat -nr
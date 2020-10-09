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
		fmt.Printf("Interface: %v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println("InterfaceByName:", err)
			return
		}

		addrs, err := byName.Addrs()
		if err != nil {
			fmt.Println("Addrs:", err)
			return
		}
		for i, addr := range addrs {
			fmt.Printf("Interface addr: %d  %s\n", i, addr.String())
		}
		//fmt.Println()
	}
}

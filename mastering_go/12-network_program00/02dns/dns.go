package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func getAddrNames(addr string) ([]string, error) {
	names, err := net.LookupAddr(addr)
	if err != nil {
		fmt.Println("Error on LookupAddr:", err)
		return nil, err
	}
	return names, nil
}

func getHostAddrs(host string) ([]string, error) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println("Error on LookupHost:", err)
		return nil, err
	}
	return addrs, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Ip Addr>\n", filepath.Base(args[0]))
		return
	}

	ipAddrStr := args[1]
	ipAddrName := net.ParseIP(ipAddrStr)
	if ipAddrName != nil {
		names, err := getAddrNames(ipAddrStr)
		if err != nil {
			fmt.Println("GetAddrNames:", err)
			return
		}
		for _, name := range names {
			fmt.Println("HostName:", name)
		}
	} else {
		addrs, err := getHostAddrs(ipAddrStr)
		if err != nil {
			fmt.Println("GetHostAddr:", err)
			return
		}
		fmt.Println("HostAddr:", addrs)
	}
}

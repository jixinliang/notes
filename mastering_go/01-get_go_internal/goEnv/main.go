package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func runtimeFunc() {
	fmt.Println("runtime compiler: ", runtime.Compiler)
	fmt.Println("goArch: ", runtime.GOARCH)
	fmt.Println("Version: ", runtime.Version())
	fmt.Println("num of cpu: ", runtime.NumCPU())
	fmt.Println("num of goroutine: ", runtime.NumGoroutine())
}

func getVersion() {
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)
	if m1 == 1 && m2 < 15 {
		fmt.Println("Need go Version 1.15 or higher!")
	}
	fmt.Println("you are using go Version 1.15 or higher!")
	fmt.Println(m1, m2)
}

func main() {
	getVersion()
}

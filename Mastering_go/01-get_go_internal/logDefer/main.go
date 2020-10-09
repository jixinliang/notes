package main

import (
	"fmt"
	"log"
	"os"
)

const LogFile = "/tmp/test.log"

func funcOne(logger *log.Logger) {
	logger.Println("---- func one ----")
	defer logger.Println("---- defer func one ----")
	for i := 0; i < 10; i++ {
		logger.Println(i)
	}
}

func funcTwo(logger *log.Logger) {
	logger.Println("---- func two ----")
	defer logger.Println("---- defer func two ----")
	for i := 0; i < 10; i++ {
		logger.Println(i)
	}
}

func main() {
	file, err := os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	logger := log.New(file, "logDefer", log.LstdFlags)
	logger.Println("Hello ")
	logger.Println("another statement")
	funcOne(logger)
	funcTwo(logger)
}

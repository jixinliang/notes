package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handler(signal os.Signal) {
	fmt.Println("Received:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		sig := <-sigs
		switch sig {
		case os.Interrupt:
			handler(sig)
		case syscall.SIGTERM:
			handler(sig)
			os.Exit(1)
		case syscall.SIGUSR2:
			fmt.Println("Handling yscall.SIGUSR2")
		default:
			fmt.Println("Ignoring:", sig)
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(time.Second * 20)
	}
}

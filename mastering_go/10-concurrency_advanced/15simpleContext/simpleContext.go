package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func f1(t int) {
	// initialize an empty context
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	go func() {
		time.Sleep(time.Second * 4)
		cancelFunc()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("f() Error:", ctx.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}

	return
}

func f2(t int) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(t)*time.Second)
	defer cancelFunc()

	go func() {
		time.Sleep(time.Second * 4)
		cancelFunc()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("f2() Error:", ctx.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}

	return
}

func f3(t int) {
	ctx := context.Background()

	deadline := time.Now().Add(time.Duration(t*2) * time.Second)
	ctx, cancelFunc := context.WithDeadline(ctx, deadline)
	defer cancelFunc()

	go func() {
		time.Sleep(time.Second * 4)
		cancelFunc()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("f3() Error:", ctx.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}

	return
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: go run %s <Number>\n", filepath.Base(args[0]))
		return
	}

	delay, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Atoi:", err)
		return
	}
	fmt.Println("Delay:", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}

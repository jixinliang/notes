package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var (
	urlStr string
	delay  int = 5
	wg     sync.WaitGroup
)

type Data struct {
	Res *http.Response
	Err error
}

func connect(ctx context.Context) error {
	defer wg.Done()
	data := make(chan Data, 1)

	transport := &http.Transport{}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Println("Request:", err)
		return err
	}

	go func() {
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Client Request:", err)
			return
		}
		pack := Data{res, err}
		data <- pack
	}()

	select {
	case <-ctx.Done():
		transport.CancelRequest(req)
		<-data
		fmt.Println("The request was cancelled")
		return ctx.Err()
	case ok := <-data:
		err := ok.Err
		res := ok.Res
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer res.Body.Close()
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Read all:", err)
			return err
		}
		fmt.Println("Server response:", string(bytes))
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <http://example.com> [Number] for delay, default 5s\n", filepath.Base(args[0]))
		return
	}

	urlStr = args[1]

	if len(args) == 3 {
		t, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Atoi:", err)
			return
		}
		delay = t
	}

	fmt.Println("Delay:", delay)

	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(delay)*time.Second)
	defer cancelFunc()

	fmt.Println("Connecting to:", urlStr)
	wg.Add(1)
	go connect(ctx)
	wg.Wait()
	fmt.Println("Exiting...")
}

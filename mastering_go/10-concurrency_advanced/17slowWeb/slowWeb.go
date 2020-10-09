package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func handler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)

	fmt.Fprintf(w, "Server url: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Delay: %d\n", delay)
	fmt.Fprintf(w, "Server Hose: %s\n", r.Host)
}

func main() {
	rand.Seed(time.Now().Unix())

	port := ":8000"
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Using default Port:", port)
	} else {
		port = ":" + args[1]
	}

	fmt.Println("Http runing...")
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}

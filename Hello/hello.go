package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s\n", time.Now())
}

func main() {
	fmt.Println("Http Running...")
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
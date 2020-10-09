package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func CheckStatusOk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Fine!")
}

func CheckNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving %s for %s\n", r.URL.Path, r.Host)
}

func main() {
	args := os.Args
	port := ":8080"

	if len(args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = ":" + args[1]
		fmt.Println("Using port:", port)
	}

	http.HandleFunc("/ok", CheckStatusOk)
	http.HandleFunc("/notfound", CheckNotFound)
	http.HandleFunc("/", MyHandler)

	fmt.Println("Http runing...")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}

// go test httpHeader* -v --count=1
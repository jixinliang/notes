package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving %s for %s\n", r.URL.Path, r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	timeStr := time.Now().Format(time.RFC1123)
	msg := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", msg)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", timeStr)
	fmt.Fprintf(w, "Serving %s for %s\n", r.URL.Path, r.Host)
}

func main() {
	args := os.Args
	port := ":8080"
	if len(args) == 1 {
		fmt.Println("Using default prot", port)
	} else {
		port = ":" + args[1]
	}

	fmt.Println("Http runing...")
	mux := http.NewServeMux()
	srv := &http.Server{Addr: port, Handler: mux, ReadTimeout: time.Second * 3, WriteTimeout: time.Second * 3}

	mux.HandleFunc("/", myHandler)
	mux.HandleFunc("/time", timeHandler)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving at: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.RFC1123)
	title := "The current time is:"
	fmt.Fprintf(w, "<h1 style=\"color: salmon\" align=\"center\">%s</h1>", title)
	fmt.Fprintf(w, "<h1 style=\"color: salmon\" align=\"center\">%s</h1>", now)
	fmt.Fprintf(w, "Serveing at: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	args := os.Args
	port := ":8080"
	if len(args) == 1 {
		fmt.Printf("Using default port: %s\n", port)
	} else {
		port = ":" + args[1]
		fmt.Printf("Using port: %s\n", port)
	}

	fmt.Println("Http runing...")

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", homeHandler)
	serveMux.HandleFunc("/time", timeHandler)

	serveMux.HandleFunc("/debug/pprof/", pprof.Index)
	serveMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	serveMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	serveMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	serveMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}

}

// go tool pprof http://localhost:8080/debug/pprof/profile
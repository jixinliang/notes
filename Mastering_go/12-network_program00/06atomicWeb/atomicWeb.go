package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
)

var count int32

func addHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count, 1)
}

func getCountHandler(w http.ResponseWriter, r *http.Request) {
	val := atomic.LoadInt32(&count)
	fmt.Println("Counts:", val)
	fmt.Fprintf(w, "<h1 align=\"center\" style=\"color: salmon\">%d</h1>", count)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	fmt.Println("Http runing...")
	http.HandleFunc("/", addHandler)
	http.HandleFunc("/get_count", getCountHandler)
	http.ListenAndServe(":8080", nil)
}

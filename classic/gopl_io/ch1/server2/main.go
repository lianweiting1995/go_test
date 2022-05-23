package server2

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func Server() {
	fmt.Println("server start")
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handleCount)
	http.ListenAndServe(":8888", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "path is %s", r.URL)
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "visiter count %d", count)
}

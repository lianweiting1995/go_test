package server1

import (
	"fmt"
	"net/http"
)

func Server() {
	fmt.Println("service start")
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path= %q\n", r.URL.Path)
}

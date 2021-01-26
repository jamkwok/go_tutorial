package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/hello", SayHelloMethod)
	http.ListenAndServe("0.0.0.0:8080", handler)
}

func SayHelloMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `Hello world`)
}

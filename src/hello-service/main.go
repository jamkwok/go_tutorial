package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.handleFunc("/api/hello", SayHelloMethod)
	http.ListenAndServe("0.0.0.0:8080", handler)
}

func SayHelloMethod() {
	fmt.Fprintf(w, `Hello world`)
}

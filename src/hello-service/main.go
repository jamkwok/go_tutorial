package main

import (
	"net/http"
	"encoding/json"
)

type Greeting struct {
    Greeting string
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/hello", SayHelloMethod)
	http.ListenAndServe("0.0.0.0:3000", handler)
}

func SayHelloMethod(w http.ResponseWriter, r *http.Request) {
	data := Greeting{Greeting: "Hello World!"}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}

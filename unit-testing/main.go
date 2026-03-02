package main

import (
	"fmt"
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Golang Ghana")
	})

	return mux
}

func main() {
	http.ListenAndServe(":8080", NewMux())
}

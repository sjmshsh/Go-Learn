package main

import (
	"fmt"
	"log"
	"net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Go")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)

	serve := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	if err := serve.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

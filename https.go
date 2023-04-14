package main

import (
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	w.Write([]byte("Hello"))
}

func main() {
	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(handle),
	}
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

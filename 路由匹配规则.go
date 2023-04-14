package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到Go Web首页！处理器为：indexHandler！")
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到Go Web欢迎页！处理器为：hiHandler！")
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到Go Web欢迎页！处理器为webHandler")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hi", hiHandler)
	mux.HandleFunc("/hi/web", webHandler)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

// 定义多个处理器
type handle1 struct{}

func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, handle1")
}

type handle2 struct{}

func (h2 *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, handle2")
}

func main() {
	handle1 := handle1{}
	handle2 := handle2{}
	// nil表示服务器使用的是默认的多路复用器DefaultServeMux
	server := http.Server{
		Addr:    "0.0.0.0:8085",
		Handler: nil,
	}
	// handle()函数调用的是多路复用器的DefaultServeMux.Handle()方法
	http.Handle("/handle1", &handle1)
	http.Handle("/handle2", &handle2)
	server.ListenAndServe()
}

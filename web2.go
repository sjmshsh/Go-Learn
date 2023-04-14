package main

import (
	"net/http"
)

// Refer 实现一个功能：在发送HTTP请求的时候，只有带上指定的refer参数
// 该请求才能调用成功，否则返回403状态
type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.baidu.com",
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:8080", referer)
}

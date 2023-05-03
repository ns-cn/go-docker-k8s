package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 监听80端口，响应所有请求并返回hello world
	http.ListenAndServe(":80",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello world")
		}))
}

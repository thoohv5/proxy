package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 设置处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// 指定监听地址和端口
	address := "127.0.0.1:8080"

	// 启动 HTTP 服务器
	fmt.Printf("Server is listening on %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

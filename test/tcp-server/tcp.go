package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 处理连接的业务逻辑
	fmt.Println("Accepted connection from", conn.RemoteAddr())

	// 在这里执行你的业务逻辑，例如读取/写入数据
	// ...

	// 关闭连接
	fmt.Println("Connection closed for", conn.RemoteAddr())
}

func main() {
	// 监听地址和端口
	listenAddr := "127.0.0.1:8082"

	// 创建监听器
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP Server is listening on", listenAddr)

	// 循环接受连接
	for {
		// 等待连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 启动一个 goroutine 处理连接
		go handleClient(conn)
	}
}

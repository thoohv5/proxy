package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务器地址和端口
	serverAddr := "127.0.0.1:8083"

	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// 连接成功后，你可以在这里执行你的业务逻辑，例如读取/写入数据

	fmt.Println("Connected to", serverAddr)

	// 例如，可以发送数据到服务器
	message := "Hello, TCP Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return
	}

	fmt.Println("Data sent successfully.")

	// 在这里你可以继续执行其他操作，例如读取服务器的响应等
}

package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器地址和端口
	serverAddress := "127.0.0.1:8085"

	// 解析服务器地址
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// 创建 UDP 连接
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	// 准备要发送的数据
	message := []byte("Hello, UDP Server!")

	// 发送数据
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending data to UDP server:", err)
		return
	}

	fmt.Println("Data sent to", serverAddress)
}

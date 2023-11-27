package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器地址和端口
	address := "127.0.0.1:8084"

	// 解析地址
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// 创建监听器
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP Server is listening on", address)

	// 接收和处理数据
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			return
		}

		fmt.Printf("Received %d bytes from %s: %s\n", n, addr, buffer[:n])

		// 在这里添加你的业务逻辑，处理接收到的数据
		// ...

		// 可以选择发送响应
		// conn.WriteToUDP([]byte("Response from UDP server"), addr)
	}
}

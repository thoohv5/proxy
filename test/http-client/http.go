package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 指定目标 URL
	url := "http://127.0.0.1:8081"

	// 发送 GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应内容
	fmt.Println("Response from", url, ":", string(body))
}

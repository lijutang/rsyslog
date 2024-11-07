package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 设置Rsyslog接收端的地址和端口
	address := "192.168.31.208:514" // Rsyslog接收服务器的地址和端口
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Println("连接失败:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 构造要发送的日志消息
	message := `Nov  1 17:49:16 localhost.localdomain nginx_access: 192.168.31.208 - - [01/Nov/2024:17:49:16 +0800] "GET /js/879.54e1b09f.js.map HTTP/1.1" 304 0 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36" "-"`

	// 发送消息
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("发送失败:", err)
		os.Exit(1)
	}

	fmt.Println("成功发送消息:", message)
}

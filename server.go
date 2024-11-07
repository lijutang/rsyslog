package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建一个UDP地址
	address := net.UDPAddr{
		Port: 514,                    // 监听的端口
		IP:   net.ParseIP("0.0.0.0"), // 接收所有IP的消息
	}

	// 创建UDP连接
	conn, err := net.ListenUDP("udp", &address)
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer conn.Close()

	fmt.Println("正在监听UDP端口 514...")

	buffer := make([]byte, 1024) // 创建缓冲区

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("读取失败:", err)
			return
		}
		fmt.Printf("来自 %s 的消息: %s\n", addr.String(), string(buffer[:n]))
	}
}

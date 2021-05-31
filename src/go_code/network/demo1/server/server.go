package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("server read err=", err)
			return
		}
		// 显示接收的消息
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听....")
	// 使用tcp协议,监听本地8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept() // 等待客户端连接
		if err != nil {
			fmt.Println("accept err=", err)
		} else {
			fmt.Println("accept suc con=", conn, " 客户端ip=%v", conn.RemoteAddr().String())
		}

		// 这里起一个协程为客户端服务
		go process(conn)
	}
	// fmt.Printf("listen suc=%v\n", listen)
}

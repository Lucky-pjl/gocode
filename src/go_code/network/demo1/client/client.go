package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn suc conn=", conn)
	reader := bufio.NewReader(os.Stdin) // os.Stdin代表标准输入
	for {
		// 从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')

		if strings.Trim(line, "\r\n") == "exit" {
			return
		}

		if err != nil {
			fmt.Println("read err=", err)
		}
		// 将line发送给服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn write err=", err)
		}
		fmt.Printf("发送了 %d 字节数据\n", n)
	}
}

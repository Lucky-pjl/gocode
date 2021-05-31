package main

import (
	"fmt"
	"go_code/network/project/server/model"
	"net"
	"time"
)

// 处理与客户端的连接
func process1(conn net.Conn) {
	defer conn.Close()

	// 调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端通讯出错,err =", err)
		return
	}
}

// 初始化 UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	// 初始化redis连接池
	initPool("119.45.252.206:6379", 16, 0, 300*time.Second)
	initUserDao()

	fmt.Println("服务器在8889端口监听...")
	lisetn, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen err =", err)
		return
	}
	defer lisetn.Close()

	// 等待客户端来连接
	for {
		fmt.Println("等待客户端连接服务器...")
		conn, err := lisetn.Accept()
		if err != nil {
			fmt.Println("accept err =", err)
		}

		// 一旦连接成功,启动一个协程和客户端保持通讯
		go process1(conn)
	}
}

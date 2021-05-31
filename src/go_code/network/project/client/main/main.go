package main

import (
	"fmt"
	"go_code/network/project/client/process"
)

var userId int
var userPwd string
var userName string

func menu() {

	// 接收用户输入
	var key int
	var loop = true
	for loop {
		fmt.Println("----------欢迎登陆多人聊天系统----------")
		fmt.Println("\t 1 登陆聊天室")
		fmt.Println("\t 2 注册用户")
		fmt.Println("\t 3 退出系统")
		fmt.Println("\t 请选择(1-3):")
		fmt.Println("---------------------------------------")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			processor := &process.UserProcess{}
			processor.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的昵称:")
			fmt.Scanf("%s\n", &userName)
			processor := &process.UserProcess{}
			processor.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有无,请重新输入")
		}
	}
}

func main() {
	menu()
}

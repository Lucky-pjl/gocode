package process

import (
	"encoding/json"
	"fmt"
	"go_code/network/project/common/message"
	"go_code/network/project/utils"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int,
	userPwd string, userName string) {

	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}

	// 2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.Register
	registerMes.UserId = userId
	registerMes.UserPwd = userPwd
	registerMes.UserName = userName

	// 3.序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}

	// 4.发送消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误, err =", err)
	}

	// 5.读取返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}

	// 6.反序列化
	var registerResMes message.RedisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功,请重新登录")
	} else if registerResMes.Code == 500 {
		fmt.Println(registerResMes.Error)
	}
	return
}

// 写一个函数，完成登陆校验
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	defer conn.Close()

	// 2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 3.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}
	mes.Data = string(data)

	// 4.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}

	// 5.将mes发送给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)

	// 7.处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}
	// 8.将mes.Data反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserName = loginMes.UserName
		CurUser.UserStatus = message.UserOnline

		// 显示当前在线用户列表
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UserIds {
			fmt.Println("用户id\t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		// 启动一个协程保持和服务器的通讯
		go serverProcessMes(conn)

		// 1.显示登陆成功的菜单
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

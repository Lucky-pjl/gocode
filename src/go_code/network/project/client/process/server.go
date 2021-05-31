package process

import (
	"encoding/json"
	"fmt"
	"go_code/network/project/common/message"
	"go_code/network/project/utils"
	"net"
	"os"
)

// 显示登陆成功后的解面

func ShowMenu() {
	fmt.Println("----------欢迎xxx登陆----------")
	fmt.Println("\t 1.显示用户在线列表")
	fmt.Println("\t 2.发送消息")
	fmt.Println("\t 3.信息列表")
	fmt.Println("\t 4.退出系统")
	fmt.Println("请选择(1-4):")

	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你的消息:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误")
	}
}

// 和服务器端保持通讯
func serverProcessMes(conn net.Conn) {

	// 创建一个transfer实例,不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("sPM readPkg err =", err)
			return
		}
		// fmt.Printf("mes=%v\n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notify message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notify)
			updateOnlineUsers(&notify)
		case message.SmsMesType: // 接收群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("消息无法处理")
		}
	}
}

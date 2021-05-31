package main

import (
	"fmt"
	"go_code/network/project/common/message"
	"go_code/network/project/server/process"
	"go_code/network/project/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端消息种类，调用不同函数
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登陆逻辑
		userPro := &process.UserProcess{
			Conn: this.Conn,
		}
		userPro.ServerProcessLogin(mes)
	case message.RegisterMesType:
		userPro := &process.UserProcess{
			Conn: this.Conn,
		}
		userPro.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {

	// 读客户端发送的信息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg err =", err)
			return err
		}
		// fmt.Println("mes=", mes)
		this.serverProcessMes(&mes)
	}
}

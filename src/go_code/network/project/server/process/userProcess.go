package process

import (
	"encoding/json"
	"fmt"
	"go_code/network/project/common/message"
	"go_code/network/project/server/model"
	"go_code/network/project/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int // 用户id
}

// 通知所有在线用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notify message.NotifyUserStatusMes
	notify.UserId = userId
	notify.Status = message.UserOnline

	// 序列化
	data, err := json.Marshal(notify)
	if err != nil {
		fmt.Println("json marshal err =", err)
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}

	// 发送 data
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 1.取出mes.Data 反序列化
	var registerMes message.Register
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json unmarshall err =", err)
		return
	}

	// 2.响应信息
	var resMes message.Message
	resMes.Type = message.RedisterResMesType
	var registerResMes message.RedisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		registerResMes.Code = 500
		registerResMes.Error = err.Error()
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	// 3.将loginResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}

	// 4.将data给 resMes
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}

	// 5.发送 data
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 1.从mes中取出mes.Data,并反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarshall err =", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginMesResType
	var loginResMes message.LoginResMes

	// 从redis中查询用户信息
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		loginResMes.Code = 500
		loginResMes.Error = err.Error()
	} else {
		loginResMes.Code = 200
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)                  // 添加在线用户
		this.NotifyOthersOnlineUser(loginMes.UserId) // 通知其他用户
		for key := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, key)
		}
		fmt.Println(user.UserName, " 登录成功")
	}

	// id=100，pwd=123456
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	loginResMes.Code = 200
	// } else {
	// 	loginResMes.Code = 500 // 表示不合法
	// 	loginResMes.Error = "该用户不存在,请注册"
	// }

	// 3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}

	// 4.将data给 resMes
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}

	// 5.发送 data
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

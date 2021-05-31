package process

import (
	"fmt"
	"go_code/network/project/client/model"
	"go_code/network/project/common/message"
)

// 客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser // 在用户登录成功后完成初始化

func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id := range onlineUsers {
		fmt.Println("用户id:", id)
	}
}

func updateOnlineUsers(notify *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notify.UserId]
	if !ok {
		user = &message.User{
			UserId:     notify.UserId,
			UserStatus: notify.Status,
		}
	}
	user.UserStatus = notify.Status
	onlineUsers[notify.UserId] = user

	outputOnlineUser()
}

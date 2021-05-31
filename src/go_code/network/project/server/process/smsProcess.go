package process

import (
	"encoding/json"
	"fmt"
	"go_code/network/project/common/message"
	"go_code/network/project/utils"
	"net"
)

type SmsProcess struct {
}

// 转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	json.Unmarshal([]byte(mes.Data), &smsMes)

	data, _ := json.Marshal(mes)

	// 遍历map
	for id, up := range userMgr.onlineUsers {
		if id != smsMes.UserId {
			this.SendMesToOnlineUser(data, up.Conn)
		}
	}
}

func (this *SmsProcess) SendMesToOnlineUser(data []byte, conn net.Conn) {

	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("消息转发失败")
	}
}

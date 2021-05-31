package process

import (
	"encoding/json"
	"go_code/network/project/common/message"
	"go_code/network/project/utils"
)

type SmsProcess struct {
}

// 发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	// 1.创建mes
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User = CurUser.User

	// 2.序列化
	data, err := json.Marshal(smsMes)
	mes.Data = string(data)
	data, err = json.Marshal(mes)

	// 3.发送
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	tf.WritePkg(data)
	return
}

package process

import (
	"encoding/json"
	"fmt"
	"go_code/network/project/common/message"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json unmarshal err =", err)
	}

	// 显示信息
	info := fmt.Sprintf("用户id:%d说:%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}

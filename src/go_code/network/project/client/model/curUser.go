package model

import (
	"go_code/network/project/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}

package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/network/project/common/message"
	"net"
)

// 将方法关联到结构体
type Transfer struct {
	Conn net.Conn
	Buf  [1024]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		err = errors.New("read pkg header error")
		return
	}
	// fmt.Println("读取到的buf =", this.Buf[:n])

	// 根据buf[:4] 转成uint32
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])
	// 读取消息内容
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		err = errors.New("read pkg data error")
		return
	}

	// 吧buf反序列化成 message类型
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("json unmarshal error")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	// 先发送一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:04])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err =", err)
		return
	}

	// 发送data
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err =", err)
		return
	}
	return
}

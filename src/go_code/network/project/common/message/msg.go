package message

const (
	LoginMesType            = "LoginMes"
	LoginMesResType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RedisterResMesType      = "RedisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// 定义消息
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code    int `json:"code"` // 状态码
	UserIds []int
	Error   string `json:"error"` // 返回错误信息
}

type Register struct {
	User `json:"user"` //类型就是User结构体
}

type RedisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// 用户状态变化信息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// SmsMes 发送消息
type SmsMes struct {
	Content string `json:"content"`
	User
}

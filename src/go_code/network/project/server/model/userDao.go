package model

import (
	"encoding/json"
	"fmt"
	"go_code/network/project/common/message"

	"github.com/garyburd/redigo/redis"
)

// 服务器启动后初始化一个全局的UserDao
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return userDao
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.Bytes(conn.Do("HGet", "users", id))
	if err == redis.ErrNil { // 表示在users中未找到对应id
		err = ERROES_USER_NOTEXISTS
		return
	}
	user = &message.User{}
	// 需要把res发序列化成 User实例
	json.Unmarshal(res, &user)
	if err != nil {
		fmt.Println("json unmarshal err =", err)
		return
	}
	return
}

// 完成登陆校验
func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	// fmt.Println("user:", user)
	if user.UserPwd != userPwd {
		err = ERRORS_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERRORS_USER_EXISTS
		return
	}
	// 向redis中添加新用户
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("添加用户错误 err =", err)
	}
	return
}

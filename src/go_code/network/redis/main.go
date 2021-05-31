package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func testString() {
	// 1.连接到redis
	conn, err := redis.Dial("tcp", "119.45.252.206:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	fmt.Println("conn succ...", conn)

	// 2.写入string
	_, err = conn.Do("Set", "name", "张三")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}
	defer conn.Close()
	fmt.Println("操作成功")

	// 3.从redis中读取string
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err =", err)
		return
	}
	// r是interface{} 需要转换
	fmt.Println("name =", r)
}

func testHash() {
	// 1.连接到redis
	conn, err := redis.Dial("tcp", "119.45.252.206:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "name", "tom")
	if err != nil {
		fmt.Println("hser err =", err)
		return
	}
	r, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hget err =", err)
		return
	}
	fmt.Println("hget name =", r)

}

// 定义一个全局的pool
var pool *redis.Pool

// 当启动程序时初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "119.45.252.206:6379")
		},
	}
}
func testPool() {
	// 先从pool中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom cat")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}

	// 取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err =", err)
		return
	}
	fmt.Println("name =", r)
}

func main() {
	// testString()
	// testHash()
	testPool()
}

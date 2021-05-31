package main

import (
	"fmt"
	"strconv"
	"time"
)

// 开启一个goroutine，每秒输出一次 "hello,world"
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

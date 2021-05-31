package main

import (
	"fmt"
)

func testSelect() {

	// 使用select可以解决从管道取数据的阻塞问题
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法遍历管道时，如果不关闭会阻塞导致deadlock
	for {
		select {
		case v := <-intChan:
			// 如果intChan一直没有关闭,不会一直阻塞而deadlock
			// 会继续向下一个case匹配
			fmt.Printf("从intChan中读取数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan中读取数据%s\n", v)
		default:
			fmt.Println("都取不到数据了")
			return
		}
	}
}

func main() {
	testSelect()
}

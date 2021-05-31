package main

import (
	"fmt"
)

func testChannel() {
	var intChan chan int
	intChan = make(chan int, 100)
	intChan <- 100
	intChan <- 200
	close(intChan) // 关闭后不能再写入数据

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入数据=", i)
		// time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读到数据=", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {

	// 创建两个管道
	intChan := make(chan int, 1)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

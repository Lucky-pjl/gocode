package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	// 关闭
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		flag = true
		num, ok := <-intChan
		if !ok {
			break
		}
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("协程退出")
	fmt.Println(len(primeChan))
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 放入结果
	exitChan := make(chan bool, 4)    // 标识退出

	go putNum(intChan)

	// 开启4个协程,从intChan取出数据,并判断是否为素数
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		// 从exitChan取出了4个结果
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数=", res)
	}

	fmt.Println("主线程退出")
}

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func testCPU() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
}

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	lock sync.Mutex
)

func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func testLock() {
	for i := 1; i <= 200; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second * 2)
	for k, v := range myMap {
		fmt.Printf("kye=%v,val=%v\n", k, v)
	}
}

func main() {
	// testLock()

}

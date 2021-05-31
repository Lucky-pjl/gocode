package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 声明一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让Camera实现 Usb接口的方法
type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func testInterface() {
	c := Computer{}
	phone := Phone{}
	camera := Camera{}

	c.Working(phone)
	c.Working(camera)

	fmt.Println("---------")
}

func testSort() {
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d", i),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println(heroes)
	fmt.Println("---------")
}

func assert() {
	var x interface{}
	var b float32 = 1.1
	x = b
	y, ok := x.(float32)
	if ok {
		fmt.Printf("y的类型是%T,值是%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("---------")
}

func main() {
	// testInterface()
	// testSort()
	assert()
}

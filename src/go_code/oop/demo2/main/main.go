package main

import (
	"fmt"
	"go_code/oop/demo2/model"
)

type Person struct {
	Name string
	Age  int
}

type Man struct {
	Person // 嵌套匿名结构体
}

func test() {
	stu := model.NewStudent("tom", 78.9)
	fmt.Println(*stu)
}

func extends() {
	p := Man{}
	p.Age = 10
	fmt.Println(p)
}

func main() {
	test()
	extends()
}

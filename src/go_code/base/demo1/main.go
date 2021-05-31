package main

import "fmt"

var (
	v1 = 100
	v2 = "fuck"
)

func main() {
	// 变量的使用
	// 方式一
	var i int = 10
	fmt.Println("i=", i)

	// 方式二
	var num = 10.11
	fmt.Println("num=", num)

	// 方式三
	str := "hello world"
	fmt.Println(str)

	// 一次定义多个int
	var n1, n2, n3 int
	fmt.Println("n11=", n1, "n2=", n2, "n3=", n3)

	// var n, name, cou = 100, "tom", 888
	n, name, cou := 100, "tom", 888
	fmt.Println("n=", n, "cou=", cou, "name=", name)

	fmt.Println("v1=", v1, "v2=", v2)
}

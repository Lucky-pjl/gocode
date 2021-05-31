package main

import (
	"fmt"
	"unsafe"
)

// 数据类型
func main() {
	// int uint byte rune(等价int32,表示一个Unicode码)
	var a int = 10
	var b uint = 1
	var c byte = 255
	fmt.Println("a=", a, "b=", b, "c=", c)

	fmt.Printf("a的类型为%T,占用的字节是 %d \n", a, unsafe.Sizeof(a))

	// 浮点数
	// float32 float64
	var f1 float32 = 12.3
	var f2 float64 = 15.888
	fmt.Printf("f1=%f , f2=%f\n", f1, f2)

	// 字符类型 使用byte
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1, " ", "c2=", c2)
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
	// 汉字
	var c3 int = '北'
	fmt.Printf("c3=%c\n", c3)

	// bool
	var isOk = true
	if isOk {
		fmt.Println("yes")
	}

	// string 字符串不可变
	var str = "hello" + "world"
	var str2 = `\n\n\n` // 原生字符串
	// str[0] = 'a' // 报错
	fmt.Println(str)
	fmt.Println(str2)

	fmt.Println()

	// 数据类型的转换 必须显示转换
	var i int = 100
	var n1 float32 = float32(i)
	fmt.Println("n1=", n1)
}

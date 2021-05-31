package main

import (
	"fmt"
	"strconv"
)

// string 与基本数据类型的转换
func main() {
	var num1 int = 99
	var num2 float64 = 23.123
	var b bool = true
	var myChar byte = 'h'
	var str string // 空字符串

	// 方式一
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)

	// 方式二 ： strconv
	str = strconv.FormatInt(123, 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	// 'f':代表格式 -ddd.dddd
	// 10:代表精度
	str = strconv.FormatFloat(123.123456, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	// strconv包中的Itoa方法
	str = strconv.Itoa(1234)
	fmt.Printf("str type %T str=%q\n", str, str)

	// string 转换为基本类型
	var str1 string = "true"
	var b1 bool
	// strconv.ParseBool(str1) 会返回两个值
	// 使用 _ 忽略不关心的值
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("\nb1 type %T b = %v\n", b1, b1)

	var n1 int64
	n1, _ = strconv.ParseInt("123456", 10, 64)
	fmt.Printf("n1 type %T n1 = %v\n", n1, n1)

	var f1 float64
	f1, _ = strconv.ParseFloat("123.1234", 64)
	fmt.Printf("f1 type %T f1 = %v\n", f1, f1)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func testStr() {
	// 1.输出语句
	fmt.Println("hello world!")

	var str string = "hello 哈哈哈"
	// for index, val := range str {
	// 	fmt.Printf("%d , %c\n", index, val)
	// }
	a, _ := strconv.Atoi("123")
	fmt.Printf("%v\n", a)
	fmt.Println(strings.Index(str, "哈"))
}

func testTime() {
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)
	fmt.Println(now.Local().Date())
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))

	// 2006/01/02 15:04:05
	fmt.Printf(now.Format("2006/01/02"))
}

func exception() {
	// 使用defer + recover来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {  // 捕获到错误
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	num := num1 / num2
	fmt.Println(num)
}

func main() {

	// testStr()
	// testTime()
	exception()
	fmt.Println("---")
}

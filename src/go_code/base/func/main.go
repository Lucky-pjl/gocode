package main

import (
	"fmt"
	"go_code/base/func/utils"
)

func add(n1 float64, n2 float64) float64 {
	return n1 + n2
}

// init函数
func init() {
	fmt.Println("init...")
}

// 函数
func main() {
	fmt.Printf("%.2f\n", add(12.3, 13.4))
	fmt.Printf("%.2f\n", utils.Add(11.0, 12.0))

	// 匿名函数
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	// f1就是函数类型
	f1 := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println("res3=", f1(30, 10))
}

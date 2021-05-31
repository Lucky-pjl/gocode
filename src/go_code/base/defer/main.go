package main

import "fmt"

// defer(延时机制,后面的语句会在函数结束后执行)
func sum(n1 int, n2 int) int {
	defer fmt.Println("ok n1 = ", n1)
	defer fmt.Println("ok n2 = ", n2)
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

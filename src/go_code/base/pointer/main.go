package main

import "fmt"

// 指针
func main() {

	var i int = 10
	fmt.Println("i的地址=", &i)

	var ptr *int = &i
	fmt.Println("ptr=", ptr, "*ptr=", *ptr, "&ptr=", &ptr)

}

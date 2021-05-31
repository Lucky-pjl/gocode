package main

import "fmt"

func main() {

	key := ""
	loop := true
	for {
		fmt.Println("家庭收支记账软件")
		fmt.Println("1.收支明细")
		fmt.Println("2.登记收入")
		fmt.Println("3.登记支出")
		fmt.Println("4.退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&key)
		switch key {
		case "1":
		case "2":
		case "3":
		case "4":
			loop = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("软件退出")
}

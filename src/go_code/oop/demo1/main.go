package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

// 给Cat绑定一个方法
func (cat Cat) test() {
	fmt.Println("test()", cat.Name)
}

func create() {
	// 方式一
	var cat1 Cat
	cat1.Name = "zs"
	cat1.Color = "黄色"
	fmt.Println(cat1)

	// 方式二
	cat2 := Cat{"ls", 2, "白色"}
	fmt.Println(cat2)

	// 方式三
	cat3 := new(Cat)
	cat3.Name = "ww"
	fmt.Println(*cat3)

	// 方式四
	cat4 := &Cat{}
	fmt.Println(cat4)
	fmt.Println("---------")
}

func method() {
	cat2 := Cat{"ls", 2, "白色"}
	cat2.test()
}

func main() {
	create()
	method()
}

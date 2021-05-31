package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取传入变量的type,kind,值
	// 1.先获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	// 2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	// 将 rVal 转成 interface{}
	iv := rVal.Interface()
	// 将 interface{} 转成int
	num2 := iv.(int)
	fmt.Println("num2=", num2)
	fmt.Println("---------")
}

func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	// 3.获取变量对应的Kind
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind=%v kind=%v\n", kind1, kind2)

	iv := rVal.Interface()
	fmt.Printf("iv=%v iv=%T\n", iv, iv)
	// 使用类型断言
	stu, ok := iv.(Student)
	if ok {
		fmt.Println("stu.Name=", stu.Name)
	}

	fmt.Println("---------")
}

type Student struct {
	Name string
	Age  int
}

func test() {
	// 对基本数据类型进行反射基本操作
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}

func reflect03(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind = %v\n", rVal.Kind())
	rVal.Elem().SetInt(20)
}

func test02() {
	var num int = 10
	reflect03(&num)
	fmt.Println("num=", num)
}

func main() {
	// test()
	test02()
}

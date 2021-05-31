package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

// 将结构体,map，切片进行序列化
func serial() {
	monster := Monster{
		Name:     "张三",
		Age:      500,
		Birthday: "2000-0-0",
		Sal:      8000.0,
		Skill:    "法外狂徒",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误,err=", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))
}

func testMap() {
	m := make(map[string]interface{})
	m["name"] = "张三"
	m["age"] = 30
	m["hobby"] = [2]string{"吃饭", "睡觉"}
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化错误,err=", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))
}

func unMarshal() {
	str := `{"name":"张三","age":500,"birthday":"2000-0-0","sal":8000,"skill":"法外狂徒"}`
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败,err=", err)
	}
	fmt.Println(monster)
}

func main() {
	serial()
	testMap()
	unMarshal()
}

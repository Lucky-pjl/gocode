package main

import (
	"testing"
)

// 编写一个测试样例,去测试addUpper是否正确
func TestAddUpper(t *testing.T) {
	// 调用
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10)执行错误,期望值=%v,实际值=%v", 55, res)
	}
	// 如果正确
	t.Logf("AddUpper(10) 执行正确")
}

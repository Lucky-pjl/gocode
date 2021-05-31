package main

import (
	"fmt"
)

// 循环
func main() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v × %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

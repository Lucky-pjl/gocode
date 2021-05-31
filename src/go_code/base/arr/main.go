package main

import "fmt"

// 数组的使用
func testArr() {
	var hens = [...]float64{1.1, 2.2, 3.3}
	total := 0.0
	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}
	fmt.Printf("%.2f\n", total)

	strs := [...]string{1: "tom", 0: "jack"}
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}

	// for-range结构遍历
	for index, val := range strs {
		fmt.Printf("%d %v\n", index, val)
	}

}

// 切片的使用
func testSlice() {
	var intArr = [...]int{1, 22, 33, 66, 99}
	// 方式一
	slice := intArr[1:3]
	fmt.Println(slice)

	// 方式二 : 通过make来创建切片
	var s1 []int = make([]int, 4)
	fmt.Println(s1)

	// 方式三
	var s2 = []int{1, 2, 3}
	fmt.Println(s2)
	fmt.Println("-------------")
}

func strSlice() {
	// string底层是byte数组,可以进行切片处理
	str := "hello@哈哈哈"
	slice := str[6:]
	fmt.Println(slice)

	// string不可变,下调语句报错
	// str[0] = 'z'

	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

	arr2 := []rune(str)
	arr2[0] = '中'
	str = string(arr2)
	fmt.Println(str)

	fmt.Println("-------------")
}

func bubbleSort() {
	arr := [...]int{24, 69, 80, 57, 13}
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	fmt.Println("-------------")
}

func multArr() {
	var arr [4][6]int
	fmt.Println(arr)
	fmt.Println("-------------")
}

// map的使用
func testMap() {
	var a map[string]string
	a = make(map[string]string, 10) // 为map分配数据空间
	a["k1"] = "v1"
	fmt.Println(a)

	// map2 := make(map[string]string)
	map3 := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}
	fmt.Println(map3)
	delete(map3, "k1") // 删除map中的数据项

	// map的遍历
	for k, v := range map3 {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	fmt.Println("-------------")
}

func main() {
	// testArr()
	testSlice()
	strSlice()
	bubbleSort()
	testMap()
}

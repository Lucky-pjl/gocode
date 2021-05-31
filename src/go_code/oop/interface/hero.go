package main

// 1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个Hero结构体切片
type HeroSlice []Hero

// 3.实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法决定使用什么标准进行排序
// 按hero的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

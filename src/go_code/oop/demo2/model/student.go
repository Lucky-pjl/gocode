package model

type student struct {
	Name  string
	Score float64
}

// student结构体首字母小写,通过工厂模式解决
func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}

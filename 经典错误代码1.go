type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000
}

// 错误原因：new无法给Show结构体里面的Param分配内存

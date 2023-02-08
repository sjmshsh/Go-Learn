type Student struct {
	Name string
}

var list map[string]Student

func main() {
	list = make(map[string]Student)
	student := Student{"LXY"}
	list["student"] = student

	/*
		方法一:
	*/
	tmpStudent := list["student"]
	tmpStudent.Name = "LDB"
	list["student"] = tmpStudent

	fmt.Println(list["student"])
}

// 是先做一次值拷贝，做出一个tmpStudent副本,然后修改该副本，
// 然后再次发生一次值拷贝复制回去，list["student"] = tmpStudent,
// 但是这种会在整体过程中发生2次结构体值拷贝，性能很差。

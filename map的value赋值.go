type Student struct {
	Name string
}

var list map[string]Student

func main() {
	list = make(map[string]Student)
	student := Student{"LXY"}
	list["student"] = student
  // 这里编译直接报错了
	list["student"].Name = "ZYQ"
	fmt.Println(list["student"])
}

// map[string]Student 的value是一个Student结构值，
// 所以当list["student"] = student,是一个值拷贝过程。
// 而list["student"]则是一个值引用。那么值引用的特点是只读。
// 所以对list["student"].Name = "LDB"的修改是不允许的。


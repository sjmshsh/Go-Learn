type Student struct {
	Name string
}

var list map[string]*Student

func main() {
	list = make(map[string]*Student)

	student := Student{"LXY"}

	list["student"] = &student
	list["student"].Name = "LDB"
	
	fmt.Println(list["student"])
}

// 我们把这个map的value修改成指针，这样指针本身是不可以修改指向的，是只读属性
// 但是指针指向的空间里面的值是可以随意修改的

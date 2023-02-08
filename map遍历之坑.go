func main() {
	// 定义map
	m := make(map[string]*student)

	// 定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 将数组依次添加到map中
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	// 打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}

// wang => wang
// zhou => wang
// li => wang  

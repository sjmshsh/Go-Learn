func main() {
	fmt.Println("hello world")
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"

	for k, v := range team {
		fmt.Println(k, v)
	}

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])
	//打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//仅打印了元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

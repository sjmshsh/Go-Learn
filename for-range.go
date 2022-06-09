func main() {
	//定义一个字符串
	var str string = "I love you"
	//方式一:普通for循环
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
	//方式二:for range
	for i, value := range str {
		fmt.Printf("索引为: %d, 具体的值为%c\n", i, value)
	}
}

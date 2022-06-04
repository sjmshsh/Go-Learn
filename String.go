func main() {
	//定义一个字符串
	var s1 string = "你好全面拥抱Go语言"
	fmt.Println(s1)
	//字符串是不可变的,这个不可变指得是字符串一旦定义好，其中的字符的值不能改变
	var s2 string = "abc"
	s2 = "def"
	//s[0] = 't'   这样是不允许的
	fmt.Println(s2)
	//字符串的表示形式
	//(1) 如果字符串中没用特殊字符，字符串的表示形式用双引号没用问题
	//(2) 如果字符串有特殊字符，就应该使用反引号
	//var s3 string = "I love you"
	var s4 string = `	//定义一个字符串
	var s1 string = "你好全面拥抱Go语言"
	fmt.Println(s1)
	//字符串是不可变的,这个不可变指得是字符串一旦定义好，其中的字符的值不能改变
	var s2 string = "abc"
	s2 = "def"
	//s[0] = 't'   这样是不允许的
	fmt.Println(s2)
	//字符串的表示形式
	//(1) 如果字符串中没用特殊字符，字符串的表示形式用双引号没用问题
	//(2) 如果字符串有特殊字符，就应该使用反引号
	var s3 string = "I love you"`
	fmt.Println(s4)
	// 字符串的拼接效果
	var s5 string = "abc" + "def"
	fmt.Println(s5)
	s5 += "higk"
	fmt.Println(s5)
	// Go语言换行原则,Go会自动加换行符，所以必须要把+留在后面
	var s6 = "abc" +
		"def"
	fmt.Println(s6)
}

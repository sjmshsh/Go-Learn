//这样就可以把函数定义成全局函数
var sub func(int, int) int

func main () {
	//定义匿名函数
	result := func(num1 int, num2 int) int {
		return num2 + num1
	}(10, 20)
	fmt.Println(result)

	//将匿名函数赋给一个变量，这个变量实际上就是函数类型的变量
	//sub等价于匿名函数
	sub := func(num1 int, num2 int) int {
		return num2 - num1
	}
	result1 := sub(50, 30)
	fmt.Println(result1)
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// 这里必须要加上&
	// var peo People = Student{}
	var peo People = &Student{}
	think := "love"
	fmt.Println(peo.Speak(think))
}

// 构成多态的三个条件：
// 1. 有interface接口并且里面定义有方法
// 2. 有子类重写的父类里面的方法(我这里说类只是方便大家理解，其实Go没有类这个概念)
// 3. 有父类的指针指向子类的具体对象

type S struct {
}

func f(x interface{}) {

}

func g(x *interface{}) {

}

func main() {
	s := S{}
	p := &s
	f(s)
	// 函数中func f(x interface{})的interface{}可以支持传入golang的任何类型
	// 包括指针，但是函数func g(x *interface{})只能接受*interface{}·
	g(s)
	f(p)
	g(p)
}

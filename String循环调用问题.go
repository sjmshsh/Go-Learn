type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}

// 在golang中 String() string ⽅法实际上是实现了 String 的接⼝的，该接⼝定义在 fmt/print.go 中：
// 在使⽤ fmt 包中的打印⽅法时，如果类型实现了这个接⼝，会直接调⽤。⽽题⽬中打印 p 的时候会直接调⽤ p
// 实现的 String() ⽅法，然后就产⽣了循环调⽤

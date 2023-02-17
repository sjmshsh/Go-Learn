type gfg struct {
	Prop string
}

func (f gfg) Geek1() string {
	return f.Prop
}

func (f gfg) Geek2() {

}

func main() {
	fooType := reflect.TypeOf(gfg{})
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
}

func main() {
	t := reflect.TypeOf(User{})
	//根据reflect.Type创建一个对象，得到该对象的指针，再根据指针提到reflect.Value
	value := reflect.New(t)
	value.Elem().FieldByName("Age").SetInt(10)
	//把反射类型转成go原始数据类型Call([]reflect.Value{})
	user := value.Interface().(*User)
}

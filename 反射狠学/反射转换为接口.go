func main() {
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

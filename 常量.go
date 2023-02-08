const cl = 100

var bl = 123

func main() {
	fmt.Println(&bl, bl)
  // 这里的&cl会在编译的时候出现错误
	fmt.Println(&cl, cl)
}

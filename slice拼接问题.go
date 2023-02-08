func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// 这里必要要加上...
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

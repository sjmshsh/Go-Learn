	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var p *int = nil
	Foo(p)
}

// non-empty interface

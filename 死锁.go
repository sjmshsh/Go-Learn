// 这段代码会发生什么？

var mu sync.Mutex
var chain string

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

func main() {
	chain = "main"
	A()
	fmt.Println(chain)
}

// 会panic
// fatal error: all goroutines are asleep - deadlock!

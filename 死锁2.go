var mu sync.RWMutex
var count int

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
}

func main() {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

// 会产生死锁panic，根据sync/rwmutex.go 中注释可以知道，读写锁当有一个协程在等待写锁时，
// 其他协程是不能获得读锁的，而在A和C中同一个调用链中间需要让出读锁，让写锁优先获取，而A的读锁又要求C调用完成，因此死锁。

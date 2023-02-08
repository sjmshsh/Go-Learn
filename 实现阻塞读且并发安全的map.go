// GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：

// type sp interface {
//    Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
//    Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
// }


type sp interface {
	// Out 存入key/val，如果该key读取的goroutine挂起，则唤醒
	// 此方法不会阻塞，时刻都可以立即执行并返回
	Out(key string, val interface{})
	// Rd 读取一个key，如果key不存在则阻塞，等待key存在或者超时
	Rd(key string, timeout time.Duration) interface{}
}

type entry struct {
	// 用于阻塞协程
	ch chan struct{}
	// 用于存放值
	value   interface{}
	isExist bool
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	// 获取这个key对应的值
	item, ok := m.c[key]
	// 如果之前没有这个数据，也就是没查出来，那么ok为false
	// 然后就会进入这个if分支
	if !ok {
		// 把数据进行插入，然后把isExist字段设置成已经存在就可以了
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	// 到这里说明key是存在的，我们需要更新key里面的val值
	item.value = val
	// 如果它不存在的化，会进入这个if分支
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}

// Rd 获取值
func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.RLock()
	// 如果找到了这个值，并且标记为存在的化进入分支直接返回结果
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.RUnlock()
		return e.value
	} else if !ok { // 如果没有查找到这个key对应的值的话
		m.rmx.RUnlock()
		m.rmx.Lock()
		e = &entry{
			ch:      make(chan struct{}),
			isExist: false,
		}
		m.c[key] = e
		m.rmx.Unlock()
		log.Println("携程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	} else {
		m.rmx.RUnlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	}
}

func main() {
	mapval := Map{
		c:   make(map[string]*entry),
		rmx: &sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		go func() {
			val := mapval.Rd("key", time.Second*6)
			log.Println("读取值为->", val)
		}()
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(val int) {
			mapval.Out("key", val)
		}(i)
	}
	time.Sleep(time.Second * 30)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

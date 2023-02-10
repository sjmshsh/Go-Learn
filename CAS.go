var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
}


// CAS是不需要循环调用的

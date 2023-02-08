// 字符串转成切片，会产生拷贝。严格来说，只要是发生类型强转都会发生内存拷贝。那么问题来了。

// 频繁的内存拷贝操作听起来对性能不大友好。有没有什么办法可以在字符串转成切片的时候不用发生拷贝呢？

// 那么如果想要在底层转换二者，只需要把 StringHeader 的地址强转成 SliceHeader 就行。那么go有个很强的包叫 unsafe 。

// unsafe.Pointer(&a)方法可以得到变量a的地址。
// (*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式。
// (*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切片的指针。
// 再通过 *转为指针指向的实际内容。

func main() {
	a := "aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a) /*获得变量a的地址*/)
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v", b)
}

// 因此这个代码只是做了地址的强制类型转换，没有把数据进行强制的拷贝

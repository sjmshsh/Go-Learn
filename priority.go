package main

import "fmt"

// 这段代码的作用是通过使用select和嵌套的for循环, 在处理ch2通道的值的时候, 优先处理ch1通道中的值
// 它可以保证在ch2通道有可接收的值的时候, 尽可能地先处理完ch1通道中的所有值, 再处理ch2通道中的值
func priority_select(ch1, ch2 <-chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val2 := <-ch2:
		priority:
			for {
				select {
				case val1 := <-ch1:
					fmt.Println(val1)
				default:
					break priority
				}
			}
			fmt.Println(val2)
		}
	}
}

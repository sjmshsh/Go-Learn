// 方法一
func main() {
   // 设置可同时使用CPU核数为1
   runtime.GOMAXPROCS(1)
   go func() {
      for i := 1; i < 101; i++ {
         // 奇数
         if i%2 == 1 {
            fmt.Println("线程1打印: ", i)
         }
         // 让出CPU
         runtime.Gosched()
      }
   }()
   go func() {
      for i := 1; i < 101; i++ {
         // 偶数
         if i%2 == 0 {
            fmt.Println("线程2打印: ", i)
         }
         // 让出CPU
         runtime.Gosched()
      }
   }()
   time.Sleep(3 * time.Second)
}

// 方法二
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		// 如果channel中没有数据，<-ch阻塞，直到channel中有一个数据时，取出
		// channel的数据并判断i的值是否是偶数，如果是偶数，则打印
		for i := 1; i < 10; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	// 第一次循环时，channel为空，可以向channel写入数据，此时i=1，所以，可以打印出1
	// 第二次循环时，channel不为空，所以 ch <- 0语句阻塞，直到channel内没有数据
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			ch <- 0
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()
	
	wg.Wait()
}

// 方法三:
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			val, _ := <-ch
			if val > 0 && val%2 == 0 {
				fmt.Println(val)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println(i)
			}
			ch <- i
		}
	}()

	wg.Wait()
}

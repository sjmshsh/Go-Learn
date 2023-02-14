func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 创建在goroutine中使用的副本
		cCp := c.Copy()
		go func() {
			// 用time.Sleep()模拟一个长任务
			time.Sleep(5 * time.Second)

			// 请注意你使用的是复制的上下文"cCp"
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()

		r.GET("/long_sync", func(c *gin.Context) {
			time.Sleep(5 * time.Second)

			log.Println("Done! in path " +
				c.Request.URL.Path)
		})
	})

	r.Run(":8080")
}

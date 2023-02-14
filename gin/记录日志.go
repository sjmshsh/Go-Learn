func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()

	// 记录到文件
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果同时需要将日志吸入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run(":8080")
}

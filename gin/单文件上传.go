func main() {
	router := gin.Default()
	// 位multipart forms设置较低的内存限制(默认是32MB)
	router.MaxMultipartMemory = 8 << 20 // 8MB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dst := "./" + file.Filename
		// 上传文件至指定完整文件路径
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("%s, uploaded!", file.Filename))
	})
	router.Run(":8080")
}

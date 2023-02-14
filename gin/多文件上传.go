func main() {
	router := gin.Default()
	// 位multipart forms设置较低的内存限制(默认是32MB)
	router.MaxMultipartMemory = 8 << 20 // 8MB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart Form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件至指定目录
			dst := "./" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	})
	router.Run(":8080")
}

// 请求如何发
// curl -X POST http://localhost:8080/upload \
//   -F "upload[]=@/Users/appleboy/test1.zip" \
//   -F "upload[]=@/Users/appleboy/test2.zip" \
//   -H "Content-Type: multipart/form-data"

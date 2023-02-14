func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(http.StatusOK, names)
	})
	router.Run(":8080")
}


// 例如这样
// while(1);[
//    "lena",
//    "austin",
//    "foo"
//]

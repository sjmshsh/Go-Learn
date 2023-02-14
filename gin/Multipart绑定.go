type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(http.StatusOK, gin.H{
					"status": "you are logged in",
				})
			} else {
				c.JSON(401, gin.H{
					"status": "unauthorized",
				})
			}
		}
	})
	router.Run(":8080")
}

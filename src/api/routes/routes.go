package routes

import "github.com/gin-gonic/gin"

func Init(router *gin.RouterGroup) {
	v1 := router.Group("v1")
	{
		// /api/v1
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{})
		})

		// /api/v1/users
		users := v1.Group("/users")
		{
			users.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{})
			})
			users.GET("/:user_id", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"user_id": c.Param("user_id"),
				})
			})
			users.POST("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"name":  c.PostForm("name"),
					"email": c.PostForm("email"),
				})
			})
			users.PUT("/:user_id", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"user_id": c.Param("user_id"),
					"name":    c.PostForm("name"),
					"email":   c.PostForm("email"),
				})
			})
			users.DELETE("/:user_id", func(c *gin.Context) {
				c.Status(204)
			})
		}
	}
}

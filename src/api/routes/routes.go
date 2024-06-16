package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	v1 := router.Group("v1")
	{
		// /api/v1
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{})
		})

		// /api/v1/users
		UsersRoutes(v1)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	usersController "github.com/marcelofabianov/go-mvc/src/api/controller"
)

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
			users.GET("/", usersController.GetUsers)
			users.GET("/:user_id", usersController.FindUserById)
			users.POST("/", usersController.CreateUser)
			users.PUT("/:user_id", usersController.UpdateUser)
			users.DELETE("/:user_id", usersController.DeleteUser)
		}
	}
}

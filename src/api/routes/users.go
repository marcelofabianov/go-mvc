package routes

import (
	"github.com/gin-gonic/gin"
	usersController "github.com/marcelofabianov/go-mvc/src/api/controller"
)

func UsersRoutes(v1 *gin.RouterGroup) {
	users := v1.Group("/users")
	{
		users.GET("/", usersController.GetUsers)
		users.GET("/:user_id", usersController.FindUserById)
		users.POST("/", usersController.CreateUser)
		users.PUT("/:user_id", usersController.UpdateUser)
		users.DELETE("/:user_id", usersController.DeleteUser)
	}
}

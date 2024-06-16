package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelofabianov/go-mvc/src/api/request"
	logger "github.com/marcelofabianov/go-mvc/src/log"
)

func FindUserById(c *gin.Context) {
	c.JSON(200, gin.H{
		"user_id": c.Param("user_id"),
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := request.ValidateRequest(err)

		logger.Error(errRest.Message, err)

		c.JSON(errRest.Status, errRest)

		return
	}

	c.JSON(201, gin.H{
		"name":  userRequest.Name,
		"email": userRequest.Email,
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"user_id": c.Param("user_id"),
		"name":    c.PostForm("name"),
		"email":   c.PostForm("email"),
	})
}

func DeleteUser(c *gin.Context) {
	c.Status(204)
}

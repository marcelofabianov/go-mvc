package usersController

import "github.com/gin-gonic/gin"

func FindUserById(c *gin.Context) {
	c.JSON(200, gin.H{
		"user_id": c.Param("user_id"),
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"name":  c.PostForm("name"),
		"email": c.PostForm("email"),
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

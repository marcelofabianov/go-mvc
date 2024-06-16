package middleware

import (
	"github.com/gin-gonic/gin"
	rest_err "github.com/marcelofabianov/go-mvc/src/errors"
)

func MethodNotAllowedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedMethods := map[string]bool{
			"GET":    true,
			"POST":   true,
			"PUT":    true,
			"DELETE": true,
			"PATCH":  true,
		}
		if !allowedMethods[c.Request.Method] {
			err := rest_err.NewMethodNotAllowedError("Method not allowed")
			c.JSON(err.Status, err)
			c.Abort()
			return
		}
		c.Next()
	}
}

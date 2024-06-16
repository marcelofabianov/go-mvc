package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marcelofabianov/go-mvc/src/api/middleware"
	"github.com/marcelofabianov/go-mvc/src/api/routes"
	rest_err "github.com/marcelofabianov/go-mvc/src/errors"
)

func main() {
	router := gin.Default()

	router.Use(middleware.MethodNotAllowedMiddleware())

	routes.Init(router.Group("/api"))

	router.NoRoute(func(c *gin.Context) {
		err := rest_err.NewNotFoundError("Resource not found")
		c.JSON(err.Status, err)
	})

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}

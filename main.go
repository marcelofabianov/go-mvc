package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marcelofabianov/go-mvc/src/api/middleware"
	"github.com/marcelofabianov/go-mvc/src/api/routes"
	rest_err "github.com/marcelofabianov/go-mvc/src/errors"
	logger "github.com/marcelofabianov/go-mvc/src/log"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.Use(middleware.MethodNotAllowedMiddleware())

	routes.Init(router.Group("/api"))

	router.NoRoute(func(c *gin.Context) {
		err := rest_err.NewNotFoundError("Resource not found")
		c.JSON(err.Status, err)
	})

	addrApi := os.Getenv("API_URL")

	logger.Info("Starting the application..." + addrApi)

	server := &http.Server{
		Addr:           addrApi,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}

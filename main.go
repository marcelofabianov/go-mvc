package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marcelofabianov/go-mvc/src/api/routes"
)

func main() {
	router := gin.Default()

	routes.Init(router.Group("/api"))

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context) {
	name := c.DefaultQuery("name", "Anonymous")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello, %s!", name),
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", ping)
		api.GET("/hello", hello)
	}

	return router
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := SetupRouter()

	r.Run(":" + port)
}

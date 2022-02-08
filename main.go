package main

import (
	"project/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.Home)
	router.POST("")
	router.Run("localhost:8080")
}

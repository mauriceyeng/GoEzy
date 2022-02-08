package main

import (
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.IndentedJSON(200, controllers.Welcome)
}

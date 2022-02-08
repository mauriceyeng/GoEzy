package main

import (
	"github/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOk, controllers.Welcome)
}

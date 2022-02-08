package main

import (
	"GoEzy/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	message := controllers.Welcome(c)
	c.IndentedJSON(200, message)
}

func RegisterUser(c *gin.Context) {
	user := controllers.UserRegister(c)
	c.IndentedJSON(http.StatusCreated, user)
}

/*To GET: curl http://localhost:8080/user/register
 */

package main

import (
	"GoEzy/structs"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) structs.User {
	newUser := structs.User{
		FirstName:     "Tony",
		LastName:      "Stark",
		Email:         "tonystark420@gmail.com",
		Contact:       "9879876452",
		City:          "go land",
		ID:            "69",
		WalletBalance: 890.2,
	}
	return newUser
}

package main

import (
	"github.com/gin-gonic/gin"
)

func newMerchant(c *gin.Context) {

	var newMerchant models.Merchant
	if err := c.BindJSON(&newMercahnt); err != nil {
		return
	} else {
		MerchantList = append(Merchant)
	}

	return newMerchant
}
func MerchantLogout(c *gin.Context) {

}

func MerchantUpload(c *gin.Context)

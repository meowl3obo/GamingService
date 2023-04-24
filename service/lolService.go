package service

import (
	. "gaming-service/model"
	. "gaming-service/config"
	provider "gaming-service/provider"

	"github.com/gin-gonic/gin"
)

func GetUserByName(c *gin.Context) {
	local := c.Param("local")
	name := c.Query("name")

	if !isCorrectCountry(local) {
		errObj := ErrorResponse {
			Code: 404, 
			Message: "查無該國家",
		}
		c.JSON(404, errObj)
		return
	}
	
	userData, statusCode, errObj := provider.GetUserByName(c, local, name)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, userData)
	}
}

func isCorrectCountry(local string) bool {
	isCorrect := false
	for _, country := range CountryList {
		if local == country {
			isCorrect = true
		}
	}
	return isCorrect
}
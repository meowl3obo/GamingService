package service

import (
	provider "gaming-service/provider"

	"github.com/gin-gonic/gin"
)

func GetUserByName(c *gin.Context) {
	local := c.Param("local")
	name := c.Query("name")
	
	userData, statusCode, errObj := provider.GetUserByName(c, local, name)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, userData)
	}
}
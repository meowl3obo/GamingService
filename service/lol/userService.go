package lolService

import (
	provider "gaming-service/provider/lol"

	"github.com/gin-gonic/gin"
)

func GetUserByName(c *gin.Context) {
	local := c.GetString("country")
	name := c.Param("name")

	userData, statusCode, errObj := provider.GetUserByName(local, name)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, userData)
	}
}

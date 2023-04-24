package service

import (
	"fmt"
	"os"
	"net/http"

	provider "gaming-service/provider"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	version := os.Getenv("VERSION")
	env := os.Getenv("ENV")
	
	c.String(http.StatusOK, fmt.Sprintf("V:%v ENV:%v", version, env))
}

func GetUserByName(c *gin.Context) {
	local := c.Param("local")
	name := c.Query("name")
	userData, statusCode, err := provider.GetUserByName(c, local, name)

	if statusCode != 200 {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, userData)
}
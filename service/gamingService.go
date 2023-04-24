package service

import (
	"fmt"
	"os"
	"net/http"

	. "gaming-service/model"
	. "gaming-service/config"
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

	if !isCorrectCountry(local) {
		errObg := ErrorResponse {
			Code: 404, 
			Message: "查無該國家",
		}
		c.JSON(404, errObg)
		return
	}
	
	userData, statusCode, err := provider.GetUserByName(c, local, name)

	if statusCode != 200 {
		errObg := ErrorResponse {
			Code: statusCode, 
			Message: err.Error(),
		}
		c.JSON(statusCode, errObg)
		return
	}
	c.JSON(statusCode, userData)
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
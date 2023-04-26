package middleware

import (
	"net/http"
	"strconv"

	. "gaming-service/config"
	. "gaming-service/model"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func CountryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		local := c.Param("local")
		if !isCorrectCountry(local) {
			errObj := ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "查無該國家",
			}
			c.JSON(http.StatusNotFound, errObj)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}

func RegionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		local := c.Param("local")
		region := CountryMap[local]
		if region == "" {
			errObj := ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "查無該國家",
			}
			c.JSON(http.StatusNotFound, errObj)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}

func CountHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		count := c.Query("count")
		_, err := strconv.Atoi(count)
		if err != nil {
			errObj := ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "count 不為數字",
			}
			c.JSON(http.StatusBadRequest, errObj)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Next()
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

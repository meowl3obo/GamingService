package middleware

import (
	"net/http"
	"strconv"

	. "gaming-service/config"
	. "gaming-service/model"

	"github.com/gin-gonic/gin"
)

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

func isCorrectCountry(local string) bool {
	isCorrect := false
	for _, country := range CountryList {
		if local == country {
			isCorrect = true
		}
	}
	return isCorrect
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
		if count != "" {
			intCount, err := strconv.Atoi(count)
			if err != nil {
				errObj := ErrorResponse{
					Code:    http.StatusBadRequest,
					Message: "count 不為數字",
				}
				c.JSON(http.StatusBadRequest, errObj)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			if intCount > 15 {
				errObj := ErrorResponse{
					Code:    http.StatusBadRequest,
					Message: "count 不能超過15",
				}
				c.JSON(http.StatusBadRequest, errObj)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}
		c.Next()
	}
}

func LangHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Query("lang")
		if lang != "" && !isCorrectLang(lang) {
			errObj := ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "查無該語系資料",
			}
			c.JSON(http.StatusNotFound, errObj)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}

func isCorrectLang(lang string) bool {
	isCorrect := false
	for _, correctLang := range LangList {
		if lang == correctLang {
			isCorrect = true
		}
	}
	return isCorrect
}

func VersionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		version := c.Query("version")
		if version != "" && !isCorrectVersion(version) {
			errObj := ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "查無該版本資料",
			}
			c.JSON(http.StatusNotFound, errObj)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}

func isCorrectVersion(version string) bool {
	isCorrect := false
	for _, correctVersion := range VersionList {
		if version == correctVersion {
			isCorrect = true
		}
	}
	return isCorrect
}

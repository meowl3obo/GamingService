package service

import (
	"fmt"
	"net/http"
	"os"

	. "gaming-service/model"

	"github.com/gin-gonic/gin"
)

// @Summary Get Version
// @Tags Version
// @version 1.0
// @produce text/plain
// @Success 200 string string 成功後返回的值
// @Router /api/version [get]
func Version(c *gin.Context) {
	version := os.Getenv("VERSION")
	env := os.Getenv("ENV")

	c.String(http.StatusOK, fmt.Sprintf("V:%v ENV:%v", version, env))
}
func NotFound(c *gin.Context) {
	error := ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "查無此路徑",
	}
	c.JSON(http.StatusNotFound, error)
}

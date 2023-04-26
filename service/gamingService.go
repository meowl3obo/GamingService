package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	version := os.Getenv("VERSION")
	env := os.Getenv("ENV")

	c.String(http.StatusOK, fmt.Sprintf("V:%v ENV:%v", version, env))
}

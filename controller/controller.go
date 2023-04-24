package controller

import (
	"gaming-service/middleware"
	"gaming-service/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Controller struct {
	*gin.Engine
}

func NewController(e *gin.Engine) *Controller {
	return &Controller{e}
}

func (r *Controller) Router() {
	r.Use(middleware.CorsMiddleware())
	r.Use(cors.Default())
	api := r.Group("/api")
	{
		api.GET("/version", service.Version)
	}
}
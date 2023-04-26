package controller

import (
	"gaming-service/middleware"
	"gaming-service/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gin.Engine
}

func NewController(e *gin.Engine) *Controller {
	return &Controller{e}
}

func (r *Controller) Router() {
	r.Use(
		middleware.CorsMiddleware(),
		cors.Default(),
	)
	api := r.Group("/api")
	{
		api.GET("/version", service.Version)

		lol := api.Group("/lol", middleware.CountryHandler())
		{
			lol.GET("/:local/user/byname", service.GetUserByName)
			lol.GET("/:local/:puuid/games", middleware.RegionHandler(), middleware.CountHandler(), service.GetGamesByPuuid)
			lol.GET("/:local/game/:matchID", middleware.RegionHandler(), service.GetGameByMatchID)
		}
	}
}

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

		lol := api.Group("/lol")
		{
			localLol := lol.Group("/:local", middleware.CountryHandler())
			{
				localLol.GET("/user/byname", service.GetUserByName)
				localLol.GET("/:puuid/games", middleware.RegionHandler(), middleware.CountHandler(), service.GetGamesByPuuid)
				localLol.GET("/game/:matchID", middleware.RegionHandler(), service.GetGameInfo)
				localLol.GET("/game/:matchID/timeline", middleware.RegionHandler(), service.GetGameTimeLine)
			}
			assetsLol := lol.Group("/assets", middleware.LangHandler(), middleware.VersionHandler())
			{
				assetsLol.GET("/roles", service.GetRoles)
			}
		}
	}
}

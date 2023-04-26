package controller

import (
	"gaming-service/middleware"
	"gaming-service/service"
	lolService "gaming-service/service/lol"

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
				localLol.GET("/:puuid/matchs", middleware.RegionHandler(), middleware.CountHandler(), lolService.GetMatchsByPuuid)
				userLol := localLol.Group("/user")
				{
					userLol.GET("/base/:name", lolService.GetUserByName)
				}
				matchLol := localLol.Group("/match")
				{
					matchLol.GET("/:matchID", middleware.RegionHandler(), lolService.GetMatchInfo)
					matchLol.GET("/:matchID/timeline", middleware.RegionHandler(), lolService.GetMatchTimeLine)
				}
			}
			assetsLol := lol.Group("/assets", middleware.LangHandler(), middleware.VersionHandler())
			{
				assetsLol.GET("/roles", lolService.GetRoles)
				assetsLol.GET("/items", lolService.GetItems)
			}
		}
	}
}

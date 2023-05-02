package controller

import (
	. "gaming-service/config"
	_ "gaming-service/docs"
	"gaming-service/middleware"
	"gaming-service/service"
	lolService "gaming-service/service/lol"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller struct {
	*gin.Engine
}

type Group struct {
	*gin.RouterGroup
}

func NewController(e *gin.Engine) *Controller {
	return &Controller{e}
}

func (r *Controller) Router() {
	r.NoRoute(service.NotFound)
	r.Use(
		middleware.CorsMiddleware(),
		cors.Default(),
	)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	api := r.Group("/api")
	{
		api.GET("/version", service.Version)

		lol := api.Group("/lol")
		{
			for key, country := range CountryMap {
				group := lol.Group(key, middleware.CountryHandler(country))
				{
					countryGroup := &Group{group}
					countryGroup.countryRoute()
				}
			}
			assetsLol := lol.Group("/assets", middleware.LangHandler(), middleware.VersionHandler())
			{
				assetsLol.GET("/roles", lolService.GetRoles)
				assetsLol.GET("/role/:name", lolService.GetRole)
				assetsLol.GET("/items", lolService.GetItems)
				assetsLol.GET("/summoners", lolService.GetSummoners)
			}
		}
	}
}

func (group *Group) countryRoute() {
	userGroup := group.Group("/user")
	{
		puuidGroup := userGroup.Group("/bypuuid")
		{
			puuidGroup.GET("/:puuid/matchs", middleware.RegionHandler(), middleware.CountHandler(), lolService.GetMatchsByPuuid)
		}
		nameGroup := userGroup.Group("/byname")
		{
			nameGroup.GET("/:name/base", lolService.GetUserByName)
		}
	}
	matchGroup := group.Group("/match")
	{
		matchGroup.GET("/:matchID", middleware.RegionHandler(), lolService.GetMatchInfo)
		matchGroup.GET("/:matchID/timeline", middleware.RegionHandler(), lolService.GetMatchTimeLine)
	}
}

package service

import (
	. "gaming-service/model"
	. "gaming-service/config"
	provider "gaming-service/provider"

	"github.com/gin-gonic/gin"
)

func GetUserByName(c *gin.Context) {
	local := c.Param("local")
	name := c.Query("name")

	userData, statusCode, errObj := provider.GetUserByName(local, name)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, userData)
	}
}

func GetGamesByPuuid(c *gin.Context) {
	local := c.Param("local")
	puuid := c.Param("puuid")
	count := c.Query("count")
	region := CountryMap[local]

	response := []MatchInfo {}

	gameIDs, statusCode, errObj := provider.GetGamesID(region, puuid, count)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	}

	for _, gameID := range gameIDs {
		gameInfo, statusCode, _ := provider.GetGameInfo(region, gameID)
		if statusCode == 200 {
			response = append(response, gameInfo)
		}
	}
	c.JSON(200, response)
}

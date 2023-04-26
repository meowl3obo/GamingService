package service

import (
	. "gaming-service/config"
	. "gaming-service/model"
	provider "gaming-service/provider"
	transfer "gaming-service/transfer"

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

	response := []MatchOverviewResponse{}

	matchIDs, statusCode, errObj := provider.GetGamesID(region, puuid, count)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	}

	for _, matchID := range matchIDs {
		gameParticipants, statusCode, _ := provider.GetGameParticipants(region, matchID)
		if statusCode == 200 {
			response = append(response, transfer.ToMatchOverviewResponse(gameParticipants))
		}
	}
	c.JSON(200, response)
}

func GetGameByMatchID(c *gin.Context) {
	local := c.Param("local")
	matchID := c.Param("matchID")
	region := CountryMap[local]

	gameTimeline, statusCode, errObj := provider.GetGameTimeline(region, matchID)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, gameTimeline)
	}
}

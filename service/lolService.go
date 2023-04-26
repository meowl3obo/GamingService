package service

import (
	. "gaming-service/config"
	. "gaming-service/model"
	provider "gaming-service/provider"
	transfer "gaming-service/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	matchOverviewChan chan MatchOverviewResponse
)

func init() {
	matchOverviewChan = make(chan MatchOverviewResponse)
}

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
		return
	}

	for _, matchID := range matchIDs {
		go getGameParticipants(region, matchID)
	}
	for len(response) < len(matchIDs) {
		data := <-matchOverviewChan
		response = append(response, data)
	}
	c.JSON(http.StatusOK, response)
}

func getGameParticipants(region string, matchID string) {
	response := MatchOverviewResponse{}
	gameParticipants, statusCode, _ := provider.GetGameParticipants(region, matchID)
	if statusCode == 200 {
		response = transfer.ToMatchOverviewResponse(gameParticipants)
	}
	matchOverviewChan <- response
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

package lolService

import (
	"fmt"
	. "gaming-service/config"
	. "gaming-service/model"
	provider "gaming-service/provider/lol"
	transfer "gaming-service/transfer"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	matchDetails chan GameInfo
)

func init() {
	matchDetails = make(chan GameInfo)
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
	forLoopCount := 0

	response := []MatchOverviewResponse{}

	matchIDs, statusCode, errObj := provider.GetGamesID(region, puuid, count)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
		return
	}

	for _, matchID := range matchIDs {
		go getGameInfo(region, matchID)
		time.Sleep(10 * time.Millisecond)
	}
	for forLoopCount < len(matchIDs) {
		forLoopCount++
		data := transfer.ToMatchOverview(<-matchDetails)
		if data.MatchID != "" {
			response = append(response, data)
		}
	}
	sort.Slice(response, func(curIndex int, nextIndex int) bool {
		return response[curIndex].StartTime > response[nextIndex].StartTime
	})
	c.JSON(http.StatusOK, response)
}

func GetGameInfo(c *gin.Context) {
	local := c.Param("local")
	matchID := c.Param("matchID")
	region := CountryMap[local]

	gameInfo, statusCode, errObj := provider.GetGameInfo(region, matchID)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, transfer.ToMatchDetails(gameInfo))
	}
}

func GetGameTimeLine(c *gin.Context) {
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

func getGameInfo(region string, matchID string) {
	response := GameInfo{}
	gameInfo, statusCode, err := provider.GetGameInfo(region, matchID)
	if statusCode == 200 {
		response = gameInfo
	} else {
		fmt.Println(region, matchID)
		fmt.Println(statusCode, err)
	}
	matchDetails <- response
}

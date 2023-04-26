package lolService

import (
	"fmt"
	. "gaming-service/config"
	. "gaming-service/model/lol"
	provider "gaming-service/provider/lol"
	transfer "gaming-service/transfer/lol"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	matchDetails chan MatchInfo
)

func init() {
	matchDetails = make(chan MatchInfo)
}

func GetMatchsByPuuid(c *gin.Context) {
	local := c.Param("local")
	puuid := c.Param("puuid")
	count := c.Query("count")
	region := CountryMap[local]
	forLoopCount := 0

	if count == "" {
		count = "10"
	}

	response := []MatchOverviewResponse{}

	matchIDs, statusCode, errObj := provider.GetMatchsID(region, puuid, count)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
		return
	}

	for _, matchID := range matchIDs {
		go getMatchInfo(region, matchID)
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

func GetMatchInfo(c *gin.Context) {
	local := c.Param("local")
	matchID := c.Param("matchID")
	region := CountryMap[local]

	matchInfo, statusCode, errObj := provider.GetMatchInfo(region, matchID)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, transfer.ToMatchDetails(matchInfo))
	}
}

func GetMatchTimeLine(c *gin.Context) {
	local := c.Param("local")
	matchID := c.Param("matchID")
	region := CountryMap[local]

	matchTimeLine, statusCode, errObj := provider.GetMatchTimeLine(region, matchID)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, matchTimeLine)
	}
}

func getMatchInfo(region string, matchID string) {
	response := MatchInfo{}
	matchInfo, statusCode, err := provider.GetMatchInfo(region, matchID)
	if statusCode == 200 {
		response = matchInfo
	} else {
		fmt.Println(region, matchID)
		fmt.Println(statusCode, err)
	}
	matchDetails <- response
}

package lolService

import (
	"net/http"
	"sort"
	"time"

	. "gaming-service/config"
	. "gaming-service/model/lol"
	provider "gaming-service/provider/lol"
	transfer "gaming-service/transfer/lol"

	"github.com/gin-gonic/gin"
)

// @Summary 取得玩家賽事列表，透過 puuid
// @Tags LoL User
// @version 1.0
// @param country path string true "國家"
// @param puuid path string true "puuid"
// @param count query string false "筆數"
// @produce application/json
// @Router /api/lol/{country}/user/bypuuid/{puuid}/matchs [get]
func GetMatchsByPuuid(c *gin.Context) {
	local := c.GetString("country")
	puuid := c.Param("puuid")
	count := c.Query("count")
	region := RegionMap[local]
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

// @Summary 取得玩家資訊，透過遊戲名稱
// @Tags LoL User
// @version 1.0
// @param country path string true "國家"
// @param name path string true "遊戲名稱"
// @produce application/json
// @Router /api/lol/{country}/user/byname/{name}/base [get]
func GetUserByName(c *gin.Context) {
	local := c.GetString("country")
	name := c.Param("name")

	userData, statusCode, errObj := provider.GetUserByName(local, name)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, userData)
	}
}

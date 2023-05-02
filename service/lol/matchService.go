package lolService

import (
	"fmt"

	. "gaming-service/config"
	. "gaming-service/model/lol"
	provider "gaming-service/provider/lol"
	transfer "gaming-service/transfer/lol"

	"github.com/gin-gonic/gin"
)

var (
	matchDetails chan MatchInfo
)

func init() {
	matchDetails = make(chan MatchInfo)
}

// @Summary 取得單賽事詳細資料
// @Tags LoL Match
// @version 1.0
// @param country path string true "國家"
// @param matchID path string true "比賽ID"
// @produce application/json
// @Router /api/lol/{country}/match/{matchID} [get]
func GetMatchInfo(c *gin.Context) {
	local := c.GetString("country")
	matchID := c.Param("matchID")
	region := RegionMap[local]

	matchInfo, statusCode, errObj := provider.GetMatchInfo(region, matchID)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, transfer.ToMatchDetails(matchInfo))
	}
}

// @Summary 取得單賽事時間線
// @Tags LoL Match
// @version 1.0
// @param country path string true "國家"
// @param matchID path string true "比賽ID"
// @produce application/json
// @Router /api/lol/{country}/match/{matchID}/timeline [get]
func GetMatchTimeLine(c *gin.Context) {
	local := c.GetString("country")
	matchID := c.Param("matchID")
	region := RegionMap[local]

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
		fmt.Println(fmt.Sprintf(
			"get match info failed\n region: %v\n matchID: %v\n statusCode: %v\n err: %v",
			region, matchID, statusCode, err))
	}
	matchDetails <- response
}

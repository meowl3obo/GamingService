package lolProvider

import (
	"errors"
	"fmt"
	"net/http"

	. "gaming-service/model"
	. "gaming-service/model/lol"
	riot "gaming-service/provider"
)

func GetMatchsID(region string, puuid string, count string) ([]string, int, ErrorResponse) {
	url := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%v/ids?start=0&count=%v", puuid, count)
	errObj := ErrorResponse{}
	res, statusCode, err := riot.RiotRequest[[]string]("GET", region, url)
	if statusCode == http.StatusNotFound {
		err = errors.New("查無該用戶")
	}
	if statusCode != http.StatusOK {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}

	return res, statusCode, errObj
}

func GetMatchInfo(region string, matchID string) (MatchInfo, int, ErrorResponse) {
	url := fmt.Sprintf("/lol/match/v5/matches/%v", matchID)
	errObj := ErrorResponse{}
	res, statusCode, err := riot.RiotRequest[MatchInfo]("GET", region, url)
	if statusCode == http.StatusNotFound {
		err = errors.New("查無該賽事")
	}
	if statusCode != http.StatusOK {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}

	return res, statusCode, errObj
}

func GetMatchTimeLine(region string, matchID string) (MatchTimeline, int, ErrorResponse) {
	url := fmt.Sprintf("/lol/match/v5/matches/%v/timeline", matchID)
	errObj := ErrorResponse{}
	res, statusCode, err := riot.RiotRequest[MatchTimeline]("GET", region, url)
	if statusCode == http.StatusNotFound {
		err = errors.New("查無該賽事")
	}
	if statusCode != http.StatusOK {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}

	return res, statusCode, errObj
}

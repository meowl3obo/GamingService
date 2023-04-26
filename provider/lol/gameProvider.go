package lolProvider

import (
	"errors"
	"fmt"
	"net/http"

	. "gaming-service/model"
	riot "gaming-service/provider"
)

func GetUserByName(local string, name string) (RiotUser, int, ErrorResponse) {
	url := fmt.Sprintf("/lol/summoner/v4/summoners/by-name/%v", name)
	errObj := ErrorResponse{}
	res, statusCode, err := riot.RiotRequest[RiotUser]("GET", local, url)
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

func GetGamesID(region string, puuid string, count string) ([]string, int, ErrorResponse) {
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

func GetGameInfo(region string, matchID string) (GameInfo, int, ErrorResponse) {
	url := fmt.Sprintf("/lol/match/v5/matches/%v", matchID)
	errObj := ErrorResponse{}
	res, statusCode, err := riot.RiotRequest[GameInfo]("GET", region, url)
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

func GetGameTimeline(region string, matchID string) (MatchTimeline, int, ErrorResponse) {
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

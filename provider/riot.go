package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	. "gaming-service/model"
)

func riotRequest[T any](method string, local string, route string) (T, int, error) {
	var response T
	header := map[string]string{
		"X-Riot-Token": os.Getenv("RIOT_TOKEN"),
	}
	baseUrl := fmt.Sprintf(os.Getenv("RIOT_API"), local)
	url := fmt.Sprintf("%v%v", baseUrl, route)
	res, err := Request(method, url, nil, header)
	if err != nil {
		return response, res.StatusCode, err
	}
	err = json.Unmarshal(res.Response, &response)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	return response, http.StatusOK, nil
}

func GetUserByName(local string, name string) (RiotUser, int, ErrorResponse) {
	url := fmt.Sprintf("/lol/summoner/v4/summoners/by-name/%v", name)
	errObj := ErrorResponse{}
	res, statusCode, err := riotRequest[RiotUser]("GET", local, url)
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
	res, statusCode, err := riotRequest[[]string]("GET", region, url)
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

func GetGameParticipants(region string, matchID string) (MatchParticipants, int, ErrorResponse) {
	url := fmt.Sprintf("/lol/match/v5/matches/%v", matchID)
	errObj := ErrorResponse{}
	res, statusCode, err := riotRequest[MatchParticipants]("GET", region, url)
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
	res, statusCode, err := riotRequest[MatchTimeline]("GET", region, url)
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


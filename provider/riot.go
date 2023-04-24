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

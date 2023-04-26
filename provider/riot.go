package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func RiotRequest[T any](method string, local string, route string) (T, int, error) {
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

func LolRequest[T any](method string, route string) (T, int, error) {
	var response T
	baseUrl := os.Getenv("LOL_ASSETS_API")
	url := fmt.Sprintf("%v%v", baseUrl, route)
	res, err := Request(method, url, nil, nil)
	if err != nil {
		if res.StatusCode == 429 {
			err = errors.New("以達到請求上限，請稍後再嘗試")
		}
		return response, res.StatusCode, err
	}
	err = json.Unmarshal(res.Response, &response)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	return response, http.StatusOK, nil
}

package provider

import (
	"fmt"
	"os"
	"encoding/json"
	"errors"

	. "gaming-service/model"

	"github.com/gin-gonic/gin"
)

func riotRequest[T any](method string, local string, route string) (T, int, error) {
	var response T
	header := map[string]string {
		"X-Riot-Token": os.Getenv("RIOT_TOKEN"),
	}
	baseUrl := fmt.Sprintf(os.Getenv("RIOT_API"), local)
	url := fmt.Sprintf("%v%v", baseUrl, route)
	res, err := Request("GET", url, nil, header)
	if err != nil {
		return response, res.StatusCode, err
	}

	err = json.Unmarshal(res.Response, &response)
	if err != nil {
		err := errors.New("反序列化失敗")
		return response, 500, err
	}

	return response, 200, nil
}

func GetUserByName(c *gin.Context, local string, name string) (RiotUser, int, error) {
	url := fmt.Sprintf("/summoner/v4/summoners/by-name/%v", name)
	res, statusCode, err := riotRequest[RiotUser]("GET", local, url)
	if statusCode == 404 {
		err = errors.New("查無該用戶")
	}
	return res, statusCode, err
}
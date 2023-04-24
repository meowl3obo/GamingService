package provider

import (
	"fmt"
	"os"
	"encoding/json"
	"errors"
	"net/http"

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
		return response, http.StatusInternalServerError, err
	}

	return response, http.StatusOK, nil
}

func GetUserByName(c *gin.Context, local string, name string) (RiotUser, int, ErrorResponse) {
	url := fmt.Sprintf("/summoner/v4/summoners/by-name/%v", name)
	errObj := ErrorResponse {}
	res, statusCode, err := riotRequest[RiotUser]("GET", local, url)
	if statusCode == http.StatusNotFound {
		err = errors.New("查無該用戶")
	}
	if statusCode != http.StatusOK {
		errObj = ErrorResponse {
			Code: statusCode, 
			Message: err.Error(),
		}
	}

	return res, statusCode, errObj
}
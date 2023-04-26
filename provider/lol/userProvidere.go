package lolProvider

import (
	"errors"
	"fmt"
	"net/http"

	. "gaming-service/model"
	. "gaming-service/model/lol"
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

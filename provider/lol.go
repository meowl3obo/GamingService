package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	. "gaming-service/model"
)

func lolRequest[T any](method string, route string) (T, int, error) {
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

func GetVersions() []string {
	route := "/api/versions.json"
	response := []string{}
	response, statusCode, err := lolRequest[[]string]("GET", route)
	if statusCode != 200 {
		fmt.Println("get version failed: ", err)
		response = append(response, "13.8.1")
	}
	return response
}

func GetLangs() []string {
	route := "/api/languages.json"
	response := []string{}
	response, statusCode, err := lolRequest[[]string]("GET", route)
	if statusCode != 200 {
		fmt.Println("get lang failed: ", err)
		response = append(response, "zh_TW", "en_US")
	}
	return response
}

func GetRolesDetails(version string, lang string) (RolesDetails, int, ErrorResponse) {
	route := fmt.Sprintf("/cdn/%v/data/%v/champion.json", version, lang)
	rolesDetails, statusCode, err := lolRequest[RolesDetails]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return rolesDetails, statusCode, errObj
}

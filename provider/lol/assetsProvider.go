package lolProvider

import (
	"fmt"

	. "gaming-service/model"
	. "gaming-service/model/lol"
	riot "gaming-service/provider"
)

func GetVersions() []string {
	route := "/api/versions.json"
	response := []string{}
	response, statusCode, err := riot.LolRequest[[]string]("GET", route)
	if statusCode != 200 {
		fmt.Println("get version failed: ", err)
		response = append(response, "13.8.1")
	}
	return response
}

func GetLangs() []string {
	route := "/api/languages.json"
	response := []string{}
	response, statusCode, err := riot.LolRequest[[]string]("GET", route)
	if statusCode != 200 {
		fmt.Println("get lang failed: ", err)
		response = append(response, "zh_TW", "en_US")
	}
	return response
}

func GetRolesDetails(version string, lang string) (RolesDetails, int, ErrorResponse) {
	route := fmt.Sprintf("/cdn/%v/data/%v/champion.json", version, lang)
	rolesDetails, statusCode, err := riot.LolRequest[RolesDetails]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return rolesDetails, statusCode, errObj
}

func GetItems(version string, lang string) (ItemsDetails, int, ErrorResponse) {
	route := fmt.Sprintf("/cdn/%v/data/%v/item.json", version, lang)
	rolesDetails, statusCode, err := riot.LolRequest[ItemsDetails]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return rolesDetails, statusCode, errObj

}

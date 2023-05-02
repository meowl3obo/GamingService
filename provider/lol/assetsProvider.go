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
	route := "/cdn/languages.json"
	response := []string{}
	response, statusCode, err := riot.LolRequest[[]string]("GET", route)
	if statusCode != 200 {
		fmt.Println("get lang failed: ", err)
		response = append(response, "zh_TW", "en_US")
	}
	return response
}

func GetRolesData(version string, lang string) (AssetsData[RoleData], int, ErrorResponse) {
	route := fmt.Sprintf("/cdn/%v/data/%v/champion.json", version, lang)
	rolesDetails, statusCode, err := riot.LolRequest[AssetsData[RoleData]]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return rolesDetails, statusCode, errObj
}

func GetRoleData(version string, lang string, name string) (AssetsData[RoleDetails], int, ErrorResponse) {
	route := fmt.Sprintf("/cdn/%v/data/%v/champion/%v.json", version, lang, name)
	roleDetails, statusCode, err := riot.LolRequest[AssetsData[RoleDetails]]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return roleDetails, statusCode, errObj
}

func GetItems(version string, lang string) (AssetsData[ItemData], int, ErrorResponse) {
	route := fmt.Sprintf("/cdn/%v/data/%v/item.json", version, lang)
	rolesDetails, statusCode, err := riot.LolRequest[AssetsData[ItemData]]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return rolesDetails, statusCode, errObj
}

func GetSummoners(version string, lang string) (AssetsData[SummonerData], int, ErrorResponse) {
	route := fmt.Sprintf("/cdn/%v/data/%v/summoner.json", version, lang)
	summonersDetails, statusCode, err := riot.LolRequest[AssetsData[SummonerData]]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return summonersDetails, statusCode, errObj
}

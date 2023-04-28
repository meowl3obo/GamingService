package lolProvider

import (
	"fmt"

	. "gaming-service/model"
	. "gaming-service/model/lol"
	riot "gaming-service/provider"
)

func GetRoleSkillDetails(mainVersion string, name string) (map[string]SkillDetail, int, ErrorResponse) {
	route := fmt.Sprintf("/%v/game/data/characters/%v/%v.bin.json", mainVersion, name, name)
	roleDetails, statusCode, err := riot.CommunityRequest[map[string]SkillDetail]("GET", route)
	errObj := ErrorResponse{}
	if err != nil {
		errObj = ErrorResponse{
			Code:    statusCode,
			Message: err.Error(),
		}
	}
	return roleDetails, statusCode, errObj
}

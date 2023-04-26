package lolTransfer

import (
	"fmt"
	. "gaming-service/model/lol"
)

func ToRoleMap(source RolesDetails) map[string]RoleResponse {
	response := map[string]RoleResponse{}

	for _, roleInfo := range source.Data {
		roleData := RoleResponse{
			Id:      roleInfo.ID,
			Key:     roleInfo.Key,
			Name:    roleInfo.Name,
			Title:   roleInfo.Title,
			Image:   fmt.Sprintf("%v/%v", roleInfo.Image.Group, roleInfo.Image.Full),
			Partype: roleInfo.Partype,
			Blurb:   roleInfo.Blurb,
			Tags:    roleInfo.Tags,
		}
		roleStatus := RoleStatus{
			Hp:                   roleInfo.Stats.Hp,
			Hpperlevel:           roleInfo.Stats.Hpperlevel,
			Mp:                   roleInfo.Stats.Mp,
			Mpperlevel:           roleInfo.Stats.Mpperlevel,
			Movespeed:            roleInfo.Stats.Movespeed,
			Armor:                roleInfo.Stats.Armor,
			Armorperlevel:        roleInfo.Stats.Armorperlevel,
			Spellblock:           roleInfo.Stats.Spellblock,
			Spellblockperlevel:   roleInfo.Stats.Spellblockperlevel,
			Attackrange:          roleInfo.Stats.Attackrange,
			Hpregen:              roleInfo.Stats.Hpregen,
			Hpregenperlevel:      roleInfo.Stats.Hpregenperlevel,
			Mpregen:              roleInfo.Stats.Mpregen,
			Mpregenperlevel:      roleInfo.Stats.Mpregenperlevel,
			Crit:                 roleInfo.Stats.Crit,
			Critperlevel:         roleInfo.Stats.Critperlevel,
			Attackdamage:         roleInfo.Stats.Attackdamage,
			Attackdamageperlevel: roleInfo.Stats.Attackdamageperlevel,
			Attackspeedperlevel:  roleInfo.Stats.Attackspeedperlevel,
			Attackspeed:          roleInfo.Stats.Attackspeed,
		}

		roleData.Status = roleStatus
		response[roleInfo.Key] = roleData
	}

	return response
}

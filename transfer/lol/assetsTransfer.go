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

func ToItemMap(source ItemsDetails) map[string]ItemResponse {
	response := map[string]ItemResponse{}

	for key, itemDetails := range source.Data {
		itemData := ItemResponse{
			Name:        itemDetails.Name,
			Description: itemDetails.Description,
			Into:        itemDetails.Into,
			From:        itemDetails.From,
			Image:       fmt.Sprintf("%v/%v", itemDetails.Image.Group, itemDetails.Image.Full),
			Tags:        itemDetails.Tags,
		}
		itemGold := ItemGold{
			Base:  itemDetails.Gold.Base,
			Total: itemDetails.Gold.Total,
			Sell:  itemDetails.Gold.Sell,
		}
		itemData.Gold = itemGold
		response[key] = itemData
	}

	return response
}

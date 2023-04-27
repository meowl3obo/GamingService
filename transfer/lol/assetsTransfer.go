package lolTransfer

import (
	"fmt"

	. "gaming-service/model/lol"
)

func ToRoleMap(source AssetsData[RoleData]) map[string]RoleResponse {
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
			Status:  setRoleStatus(roleInfo.Stats),
		}
		response[roleInfo.ID] = roleData
	}

	return response
}

func ToItemMap(source AssetsData[ItemData]) map[string]ItemResponse {
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

func ToSummonerMap(source AssetsData[SummonerData]) map[string]SummonerResponse {
	response := map[string]SummonerResponse{}
	for key, summonerDetails := range source.Data {
		summonerData := SummonerResponse{
			Id:          summonerDetails.ID,
			Name:        summonerDetails.Name,
			Description: summonerDetails.Description,
			Cooldown:    summonerDetails.Cooldown[0],
			Key:         summonerDetails.Key,
			Range:       summonerDetails.Range[0],
			Image:       fmt.Sprintf("%v/%v", summonerDetails.Image.Group, summonerDetails.Image.Full),
		}
		response[key] = summonerData
	}

	return response
}

func ToRoleDetails(source AssetsData[RoleDetails], name string) RoleDetailsResponse {
	data := source.Data[name]
	response := RoleDetailsResponse{
		Id:      data.ID,
		Key:     data.Key,
		Name:    data.Name,
		Title:   data.Title,
		Image:   fmt.Sprintf("%v/%v", data.Image.Group, data.Image.Full),
		Partype: data.Partype,
		Blurb:   data.Blurb,
		Tags:    data.Tags,
		Status:  setRoleStatus(data.Stats),
		Skill: SkillGroup{
			Passive: Skill{
				Name:        "",
				Description: "",
				Colldown:    []int{},
				Cost:        []int{},
				CostType:    "",
				Range:       []int{},
				Image:       "",
			},
			Q: Skill{
				Name:        "",
				Description: "",
				Colldown:    []int{},
				Cost:        []int{},
				CostType:    "",
				Range:       []int{},
				Image:       ""},
			W: Skill{
				Name:        "",
				Description: "",
				Colldown:    []int{},
				Cost:        []int{},
				CostType:    "",
				Range:       []int{},
				Image:       ""},
			E: Skill{
				Name:        "",
				Description: "",
				Colldown:    []int{},
				Cost:        []int{},
				CostType:    "",
				Range:       []int{},
				Image:       ""},
			R: Skill{
				Name:        "",
				Description: "",
				Colldown:    []int{},
				Cost:        []int{},
				CostType:    "",
				Range:       []int{},
				Image:       ""},
		},
	}

	return response
}

func setRoleStatus(stats RoleState) RoleStatus {
	roleStatus := RoleStatus{
		Hp:                   stats.Hp,
		Hpperlevel:           stats.Hpperlevel,
		Mp:                   stats.Mp,
		Mpperlevel:           stats.Mpperlevel,
		Movespeed:            stats.Movespeed,
		Armor:                stats.Armor,
		Armorperlevel:        stats.Armorperlevel,
		Spellblock:           stats.Spellblock,
		Spellblockperlevel:   stats.Spellblockperlevel,
		Attackrange:          stats.Attackrange,
		Hpregen:              stats.Hpregen,
		Hpregenperlevel:      stats.Hpregenperlevel,
		Mpregen:              stats.Mpregen,
		Mpregenperlevel:      stats.Mpregenperlevel,
		Crit:                 stats.Crit,
		Critperlevel:         stats.Critperlevel,
		Attackdamage:         stats.Attackdamage,
		Attackdamageperlevel: stats.Attackdamageperlevel,
		Attackspeedperlevel:  stats.Attackspeedperlevel,
		Attackspeed:          stats.Attackspeed,
	}

	return roleStatus
}

func setSkill() Skill {
	skill := Skill{}

	return skill
}

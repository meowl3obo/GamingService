package lolTransfer

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	. "gaming-service/model/lol"

	"golang.org/x/exp/slices"
)

func ToRoleMap(source AssetsData[RoleData]) map[string]RoleResponse {
	response := map[string]RoleResponse{}

	for _, roleInfo := range source.Data {
		roleData := RoleResponse{
			Id:      roleInfo.ID,
			Key:     roleInfo.Key,
			Name:    roleInfo.Name,
			Title:   roleInfo.Title,
			Image:   setImageRoute(roleInfo.Image),
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
			Image:       setImageRoute(itemDetails.Image),
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
			Image:       setImageRoute(summonerDetails.Image),
		}
		response[key] = summonerData
	}

	return response
}

func ToRoleDetails(source AssetsData[RoleDetails], name string, skillDetails map[string]SkillDetail) RoleDetailsResponse {
	data := source.Data[name]
	skillQKeys := fmt.Sprintf("Characters/%v/Spells/%vAbility/%v", name, data.Spells[0].ID, data.Spells[0].ID)
	skillWKeys := fmt.Sprintf("Characters/%v/Spells/%vAbility/%v", name, data.Spells[1].ID, data.Spells[1].ID)
	skillEKeys := fmt.Sprintf("Characters/%v/Spells/%vAbility/%v", name, data.Spells[2].ID, data.Spells[2].ID)
	skillRKeys := fmt.Sprintf("Characters/%v/Spells/%vAbility/%v", name, data.Spells[3].ID, data.Spells[3].ID)
	response := RoleDetailsResponse{
		Id:      data.ID,
		Key:     data.Key,
		Name:    data.Name,
		Title:   data.Title,
		Image:   setImageRoute(data.Image),
		Partype: data.Partype,
		Blurb:   data.Blurb,
		Tags:    data.Tags,
		Status:  setRoleStatus(data.Stats),
		Skill: SkillGroup{
			Passive: Passive{
				Name:        data.Passive.Name,
				Description: data.Passive.Description,
				Image:       setImageRoute(data.Passive.Image),
			},
			Q: setSkill(data.Spells[0], data.Partype, skillDetails[skillQKeys].MSpell.MDataValues),
			W: setSkill(data.Spells[1], data.Partype, skillDetails[skillWKeys].MSpell.MDataValues),
			E: setSkill(data.Spells[2], data.Partype, skillDetails[skillEKeys].MSpell.MDataValues),
			R: setSkill(data.Spells[3], data.Partype, skillDetails[skillRKeys].MSpell.MDataValues),
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

func setSkill(sourceSkill RoleSkill, partype string, values []SKillDataValues) Skill {
	skill := Skill{
		Name:        sourceSkill.Name,
		Description: sourceSkill.Description,
		Tooltip:     setTooltip(values, sourceSkill.Tooltip),
		Cooldown:    setCoolDown(sourceSkill.Cooldown),
		Cost:        setCostTemplate(sourceSkill.Resource, sourceSkill.Cost, partype, values),
		Range:       setRange(sourceSkill.Range),
		Image:       setImageRoute(sourceSkill.Image),
	}

	return skill
}

func setRange(rangeList []int) string {
	skillRange := setSkillValueString(rangeList)
	return skillRange
}

func setCoolDown(coolDownList []float64) string {
	skillCoolDown := setSkillValueString(coolDownList)
	return skillCoolDown
}

func setCostTemplate(template string, costList []int, costType string, values []SKillDataValues) string {
	skillCost := setSkillValueString(costList)
	template = strings.ReplaceAll(template, "{{ abilityresourcename }}", costType)
	template = strings.ReplaceAll(template, "{{ cost }}", skillCost)
	template = setTooltip(values, template)
	return template
}

func setTooltip(values []SKillDataValues, tooltip string) string {
	spellCostReg, _ := regexp.Compile(`{{ \S+ }}`)
	spellCostKeyReg, _ := regexp.Compile(`\w+`)
	spellCostAdditionReg, _ := regexp.Compile(`\d+`)
	damageReg, _ := regexp.Compile(`^\wdamage|^damage$|^totaldamage$`)
	damageEdgeReg, _ := regexp.Compile(`^\wedgedamage`)
	damageRatioReg, _ := regexp.Compile(`^\wtotal\w+ratio`)
	damageBonusReg, _ := regexp.Compile(`^\wsweetspotbonus`)
	spellCosts := spellCostReg.FindAllString(tooltip, -1)
	for _, spellCost := range spellCosts {
		checkCost := spellCostKeyReg.FindString(spellCost)
		additionString := spellCostAdditionReg.FindString(spellCost)
		if damageReg.MatchString(checkCost) || damageEdgeReg.MatchString(checkCost) {
			valueString := ""
			bouns := []float64{}
			valueIndex := slices.IndexFunc(values, func(value SKillDataValues) bool {
				return strings.Contains(strings.ToLower(value.MName), "basedamage")
			})
			ratioIndex := slices.IndexFunc(values, func(value SKillDataValues) bool {
				return damageRatioReg.MatchString(strings.ToLower(value.MName))
			})
			if damageEdgeReg.MatchString(checkCost) {
				bounsIndex := slices.IndexFunc(values, func(value SKillDataValues) bool {
					return damageBonusReg.MatchString(strings.ToLower(value.MName))
				})
				if bounsIndex > -1 {
					bouns = values[bounsIndex].MValues
				}
			}
			if valueIndex > -1 {
				value := values[valueIndex]
				valueString = fmt.Sprintf("%v%v", valueString, reduceSkill(value, additionString, bouns))
			}
			if ratioIndex > -1 {
				value := values[ratioIndex]
				valueString = fmt.Sprintf("%v + %v", valueString, reduceSkillRatio(value, bouns))
			}
			tooltip = strings.ReplaceAll(tooltip, spellCost, valueString)
		} else {
			valueIndex := slices.IndexFunc(values, func(value SKillDataValues) bool {
				return strings.ToLower(value.MName) == checkCost
			})
			if valueIndex > -1 {
				value := values[valueIndex]
				valueString := reduceSkill(value, additionString, []float64{})
				tooltip = strings.ReplaceAll(tooltip, spellCost, valueString)
			}
		}
	}
	tooltip = strings.Replace(tooltip, "{{ spellmodifierdescriptionappend }}", "", 1)
	return spellCostReg.ReplaceAllString(tooltip, "'?'")
}

func reduceSkill(value SKillDataValues, additionString string, bouns []float64) string {
	valueString := ""
	for i, skillValue := range value.MValues {
		if i != 0 && i != (len(value.MValues)-1) {
			if additionString != "" {
				addition, err := strconv.ParseFloat(additionString, 64)
				if err == nil {
					skillValue = skillValue * addition
				}
			}
			if len(bouns) > 0 {
				skillValue = skillValue + skillValue*bouns[i]
			}
			skillValue = math.Round(math.Abs(skillValue)*100) / 100
			valueString = fmt.Sprintf("%v/%v", valueString, skillValue)
		}
	}
	return strings.Replace(valueString, "/", "", 1)
}

func reduceSkillRatio(value SKillDataValues, bouns []float64) string {
	valueString := "("
	for i, skillValue := range value.MValues {
		if i != 0 && i != (len(value.MValues)-1) {
			skillValue = skillValue * 100
			if len(bouns) > 0 {
				skillValue = math.Round(math.Abs(skillValue+skillValue*bouns[i])*100) / 100
			}
			valueString = fmt.Sprintf("%v/%v%v", valueString, math.Round(skillValue*100)/100, "%")
		}
	}
	return fmt.Sprintf("%v)", strings.Replace(valueString, "/", "", 1))
}

func setSkillValueString[T any](values []T) string {
	skillValue := ""
	for _, value := range values {
		skillValue = fmt.Sprintf("%v/%v", skillValue, value)
	}
	skillValue = strings.Replace(skillValue, "/", "", 1)
	return skillValue
}

func setImageRoute(sourceImage AssetsImage) string {
	return fmt.Sprintf("%v/%v", sourceImage.Group, sourceImage.Full)
}

package lolTransfer

import (
	"strconv"

	. "gaming-service/model/lol"
)

func ToMatchOverview(source MatchInfo) MatchOverviewResponse {
	response := MatchOverviewResponse{}
	participans := []ParticipantOverview{}

	for _, team := range source.Info.Teams {
		if team.Win {
			response.WinTeam = team.TeamID
		}
	}
	for index, sourceParticipant := range source.Info.Participants {
		participan := ParticipantOverview{
			Puuid:   source.Metadata.Participants[index],
			Name:    sourceParticipant.SummonerName,
			Team:    sourceParticipant.TeamID,
			Kills:   sourceParticipant.Kills,
			Deaths:  sourceParticipant.Deaths,
			Assists: sourceParticipant.Assists,
			Role:    sourceParticipant.ChampionName,
			RoleId:  sourceParticipant.ChampionID,
		}
		participans = append(participans, participan)
	}

	response.Mode = source.Info.GameMode
	response.StartTime = source.Info.GameStartTimestamp
	response.EndTime = source.Info.GameEndTimestamp
	response.MatchID = source.Metadata.MatchID
	response.Participants = participans
	return response
}

func ToMatchDetails(source MatchInfo) MatchDetailsResponse {
	response := MatchDetailsResponse{}
	participans := []ParticipantDetails{}
	teams := map[string]TeamDetails{}

	for _, team := range source.Info.Teams {
		if team.Win {
			response.WinTeam = team.TeamID
		}
	}
	for index, sourceParticipant := range source.Info.Participants {
		participan := ParticipantDetails{
			Puuid:      source.Metadata.Participants[index],
			Name:       sourceParticipant.SummonerName,
			Team:       sourceParticipant.TeamID,
			Kills:      sourceParticipant.Kills,
			Deaths:     sourceParticipant.Deaths,
			Assists:    sourceParticipant.Assists,
			Role:       sourceParticipant.ChampionName,
			RoleId:     sourceParticipant.ChampionID,
			Level:      sourceParticipant.ChampLevel,
			FirstBlood: sourceParticipant.FirstBloodKill,
			Item: Item{
				Item0: sourceParticipant.Item0,
				Item1: sourceParticipant.Item1,
				Item2: sourceParticipant.Item2,
				Item3: sourceParticipant.Item3,
				Item4: sourceParticipant.Item4,
				Item5: sourceParticipant.Item5,
			},
		}
		participans = append(participans, participan)
	}
	for _, sourceTeam := range source.Info.Teams {
		teams[strconv.Itoa(sourceTeam.TeamID)] = TeamDetails{
			Baron:      sourceTeam.Objectives.Baron.Kills,
			Kills:      sourceTeam.Objectives.Champion.Kills,
			Dragon:     sourceTeam.Objectives.Dragon.Kills,
			RiftHerald: sourceTeam.Objectives.RiftHerald.Kills,
			Inhibitor:  sourceTeam.Objectives.Inhibitor.Kills,
			Tower:      sourceTeam.Objectives.Tower.Kills,
		}
	}

	response.Mode = source.Info.GameMode
	response.StartTime = source.Info.GameStartTimestamp
	response.EndTime = source.Info.GameEndTimestamp
	response.MatchID = source.Metadata.MatchID
	response.Participants = participans
	response.Teams = teams

	return response
}

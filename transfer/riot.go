package transfer

import (
	. "gaming-service/model"
)

func ToMatchOverviewResponse(source MatchParticipants) MatchOverviewResponse {
	response := MatchOverviewResponse{}
	participans := []ParticipantOverview{}

	response.Mode = source.Info.GameMode
	response.StartTime = source.Info.GameStartTimestamp
	response.EndTime = source.Info.GameEndTimestamp

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
		}
		participans = append(participans, participan)
	}

	response.Participants = participans
	return response
}

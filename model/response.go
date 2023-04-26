package model

type MatchOverviewResponse struct {
	Mode         string
	StartTime    int64
	EndTime      int64
	WinTeam      int
	Participants []ParticipantOverview
}

type MatchDetailsResponse struct {
	Mode         string
	StartTime    int64
	EndTime      int64
	WinTeam      int
	Participants []ParticipantDetails
}
type ParticipantOverview struct {
	Puuid   string
	Name    string
	Team    int
	Kills   int // 擊殺
	Deaths  int // 死亡
	Assists int // 助攻
}

type ParticipantDetails struct {
	Puuid   string
	Name    string
	Team    int
	Kills   int // 擊殺
	Deaths  int // 死亡
	Assists int // 助攻
}

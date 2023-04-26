package model

type MatchOverviewResponse struct {
	MatchID      string
	Mode         string
	StartTime    int64
	EndTime      int64
	WinTeam      int
	Participants []ParticipantOverview
}

type MatchDetailsResponse struct {
	MatchID      string
	Mode         string
	StartTime    int64
	EndTime      int64
	WinTeam      int
	Participants []ParticipantDetails
	Teams        map[string]TeamDetails
}

type ParticipantOverview struct {
	Puuid   string
	Name    string
	Team    int
	Kills   int // 擊殺
	Deaths  int // 死亡
	Assists int // 助攻
	Role    string
}

type ParticipantDetails struct {
	Puuid      string
	Name       string
	Team       int
	Kills      int // 擊殺
	Deaths     int // 死亡
	Assists    int // 助攻
	Role       string
	Level      int
	FirstBlood bool
	Item       Item
}

type Item struct {
	Item0 int
	Item1 int
	Item2 int
	Item3 int
	Item4 int
	Item5 int
}

type TeamDetails struct {
	Baron      int
	Kills      int
	Dragon     int
	RiftHerald int // 寓示者
	Inhibitor  int // 水晶兵營
	Tower      int
}

package lolModel

type RoleResponse struct {
	Id      string
	Key     string
	Name    string
	Title   string
	Image   string
	Partype string
	Blurb   string
	Tags    []string
	Status  RoleStatus
}

type RoleStatus struct {
	Hp                   float64 //血量
	Hpperlevel           float64 //成長血量
	Mp                   float64 //MP
	Mpperlevel           float64 //成長MP
	Movespeed            int     //移速
	Armor                int     //物防
	Armorperlevel        float64 //成長物防
	Spellblock           float64 //魔防
	Spellblockperlevel   float64 //成長魔防
	Attackrange          int     //攻擊距離
	Hpregen              float64 //生命回復
	Hpregenperlevel      float64 //成長生命回復
	Mpregen              float64 //mp回復
	Mpregenperlevel      float64 //成長mp回復
	Crit                 int     //爆擊率
	Critperlevel         int     //成長爆擊率
	Attackdamage         float64 //物攻
	Attackdamageperlevel float64 //成長物攻
	Attackspeedperlevel  float64 //成長攻速
	Attackspeed          float64 //攻速
}

type ItemResponse struct {
	Name        string
	Description string
	Into        []string
	From        []string
	Image       string
	Gold        ItemGold
	Tags        []string
}

type ItemGold struct {
	Base  int
	Total int
	Sell  int
}

type SummonerResponse struct {
	Id          string
	Name        string
	Description string
	Cooldown    int
	Key         string
	Range       int
	Image       string
}

type RoleDetailsResponse struct {
	Id      string
	Key     string
	Name    string
	Title   string
	Image   string
	Partype string
	Blurb   string
	Tags    []string
	Status  RoleStatus
	Skill   SkillGroup
}

type SkillGroup struct {
	Passive Passive
	Q       Skill
	W       Skill
	E       Skill
	R       Skill
}

type Passive struct {
	Name        string
	Description string
	Image       string
}

type Skill struct {
	Name        string
	Description string
	Tooltip     string
	Cooldown    string
	Cost        string
	Range       string
	Image       string
}

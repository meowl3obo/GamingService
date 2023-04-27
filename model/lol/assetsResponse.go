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
	Hp                   int     //血量
	Hpperlevel           int     //成長血量
	Mp                   int     //MP
	Mpperlevel           float64 //成長MP
	Movespeed            int     //移速
	Armor                int     //物防
	Armorperlevel        float64 //成長物防
	Spellblock           int     //魔防
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

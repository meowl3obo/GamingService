package lolModel

type AssetsData[T any] struct {
	Type    string       `json:"type"`
	Version string       `json:"version"`
	Data    map[string]T `json:"data"`
}

type RoleData struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Blurb string `json:"blurb"`
	Info  struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Image   AssetsImage `json:"image"`
	Tags    []string    `json:"tags"`
	Partype string      `json:"partype"`
	Stats   RoleState   `json:"stats"`
}

type ItemData struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Colloq      string      `json:"colloq"`
	Plaintext   string      `json:"plaintext"`
	Into        []string    `json:"into"`
	From        []string    `json:"from"`
	Image       AssetsImage `json:"image"`
	Gold        struct {
		Base        int  `json:"base"`
		Purchasable bool `json:"purchasable"`
		Total       int  `json:"total"`
		Sell        int  `json:"sell"`
	} `json:"gold"`
	Tags []string `json:"tags"`
	Maps struct {
		Num11 bool `json:"11"`
		Num12 bool `json:"12"`
		Num21 bool `json:"21"`
		Num22 bool `json:"22"`
	} `json:"maps"`
	Stats struct {
		FlatMovementSpeedMod int `json:"FlatMovementSpeedMod"`
	} `json:"stats"`
}

type SummonerData struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Tooltip      string `json:"tooltip"`
	Maxrank      int    `json:"maxrank"`
	Cooldown     []int  `json:"cooldown"`
	CooldownBurn string `json:"cooldownBurn"`
	Cost         []int  `json:"cost"`
	CostBurn     string `json:"costBurn"`
	Datavalues   struct {
	} `json:"datavalues"`
	Effect        []any       `json:"effect"`
	EffectBurn    []any       `json:"effectBurn"`
	Vars          []any       `json:"vars"`
	Key           string      `json:"key"`
	SummonerLevel int         `json:"summonerLevel"`
	Modes         []string    `json:"modes"`
	CostType      string      `json:"costType"`
	Maxammo       string      `json:"maxammo"`
	Range         []int       `json:"range"`
	RangeBurn     string      `json:"rangeBurn"`
	Image         AssetsImage `json:"image"`
	Resource      string      `json:"resource"`
}

type RoleDetails struct {
	ID    string      `json:"id"`
	Key   string      `json:"key"`
	Name  string      `json:"name"`
	Title string      `json:"title"`
	Image AssetsImage `json:"image"`
	Skins []struct {
		ID      string `json:"id"`
		Num     int    `json:"num"`
		Name    string `json:"name"`
		Chromas bool   `json:"chromas"`
	} `json:"skins"`
	Lore      string   `json:"lore"`
	Blurb     string   `json:"blurb"`
	Allytips  []string `json:"allytips"`
	Enemytips []string `json:"enemytips"`
	Tags      []string `json:"tags"`
	Partype   string   `json:"partype"`
	Info      struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Stats       RoleState   `json:"stats"`
	Spells      []RoleSkill `json:"spells"`
	Passive     RolePassive `json:"passive"`
	Recommended []any       `json:"recommended"`
}

type RoleState struct {
	Hp                   float64 `json:"hp"`
	Hpperlevel           float64 `json:"hpperlevel"`
	Mp                   float64 `json:"mp"`
	Mpperlevel           float64 `json:"mpperlevel"`
	Movespeed            int     `json:"movespeed"`
	Armor                int     `json:"armor"`
	Armorperlevel        float64 `json:"armorperlevel"`
	Spellblock           float64 `json:"spellblock"`
	Spellblockperlevel   float64 `json:"spellblockperlevel"`
	Attackrange          int     `json:"attackrange"`
	Hpregen              float64 `json:"hpregen"`
	Hpregenperlevel      float64 `json:"hpregenperlevel"`
	Mpregen              float64 `json:"mpregen"`
	Mpregenperlevel      float64 `json:"mpregenperlevel"`
	Crit                 int     `json:"crit"`
	Critperlevel         int     `json:"critperlevel"`
	Attackdamage         float64 `json:"attackdamage"`
	Attackdamageperlevel float64 `json:"attackdamageperlevel"`
	Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
	Attackspeed          float64 `json:"attackspeed"`
}

type RoleSkill struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tooltip     string `json:"tooltip"`
	Leveltip    struct {
		Label  []string `json:"label"`
		Effect []string `json:"effect"`
	} `json:"leveltip"`
	Maxrank      int       `json:"maxrank"`
	Cooldown     []float64 `json:"cooldown"`
	CooldownBurn string    `json:"cooldownBurn"`
	Cost         []int     `json:"cost"`
	CostBurn     string    `json:"costBurn"`
	Datavalues   struct {
	} `json:"datavalues"`
	Effect     []any       `json:"effect"`
	EffectBurn []any       `json:"effectBurn"`
	Vars       []any       `json:"vars"`
	CostType   string      `json:"costType"`
	Maxammo    string      `json:"maxammo"`
	Range      []int       `json:"range"`
	RangeBurn  string      `json:"rangeBurn"`
	Image      AssetsImage `json:"image"`
	Resource   string      `json:"resource"`
}

type RolePassive struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       AssetsImage `json:"image"`
}

type AssetsImage struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

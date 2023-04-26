package lolModel

type RolesDetails struct {
	Type    string                 `json:"type"`
	Version string                 `json:"version"`
	Data    map[string]RoleDetails `json:"data"`
}

type RoleDetails struct {
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
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   struct {
		Hp                   int     `json:"hp"`
		Hpperlevel           int     `json:"hpperlevel"`
		Mp                   int     `json:"mp"`
		Mpperlevel           float64 `json:"mpperlevel"`
		Movespeed            int     `json:"movespeed"`
		Armor                int     `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Spellblock           int     `json:"spellblock"`
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
	} `json:"stats"`
}

type ItemsDetails struct {
	Type    string                 `json:"type"`
	Version string                 `json:"version"`
	Data    map[string]ItemDetails `json:"data"`
}

type ItemDetails struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Colloq      string   `json:"colloq"`
	Plaintext   string   `json:"plaintext"`
	Into        []string `json:"into"`
	From        []string `json:"from"`
	Image       struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Gold struct {
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

package lolModel

type SkillDetail struct {
	MSpell struct {
		MDataValues []SKillDataValues `json:"mDataValues"`
		Type        string            `json:"__type"`
	} `json:"mSpell"`
}

type SKillDataValues struct {
	MName   string    `json:"mName"`
	MValues []float64 `json:"mValues"`
}

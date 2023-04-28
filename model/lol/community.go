package lolModel

type SkillDetail struct {
	MSpell struct {
		MDataValues []struct {
			MName   string    `json:"mName"`
			MValues []float64 `json:"mValues"`
		} `json:"mDataValues"`
		Type string `json:"__type"`
	} `json:"mSpell"`
}

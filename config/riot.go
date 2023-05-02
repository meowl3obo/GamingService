package config

var (
	CountryList = []string{"BR1", "EUN1", "EUW1", "JP1", "KR", "LA1", "LA2", "NA1", "OC1", "PH2", "RU", "SG2", "TH2", "TR1", "TW2", "VN2"}
	CountryMap  = map[string]string{
		"br":  "BR1",
		"eun": "EUN1",
		"euw": "EUW1",
		"jp":  "JP1",
		"kr":  "KR",
		"la1": "LA1",
		"la2": "LA2",
		"na":  "NA1",
		"oc":  "OC1",
		"ph":  "PH2",
		"ru":  "RU",
		"sh":  "SG2",
		"th":  "TH2",
		"tr":  "TR1",
		"tw":  "TW2",
		"vn":  "VN2",
	}
	RegionMap = map[string]string{
		"TW2":  "SEA",
		"OC1":  "SEA",
		"PH2":  "SEA",
		"SG2":  "SEA",
		"TH2":  "SEA",
		"VN2":  "SEA",
		"KR":   "ASIA",
		"JP1":  "ASIA",
		"EUN1": "EUROPE",
		"EUW1": "EUROPE",
		"TR1":  "EUROPE",
		"RU":   "EUROPE",
		"NA1":  "AMERICAS",
		"BR1":  "AMERICAS",
		"LA1":  "AMERICAS",
		"LA2":  "AMERICAS",
	}
	LangList    = []string{}
	VersionList = []string{}
)

func SetLangList(langList []string) {
	LangList = langList
}

func SetVersionList(versionList []string) {
	VersionList = versionList
}

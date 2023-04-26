package config

var (
	CountryList = []string{"BR1", "EUN1", "EUW1", "JP1", "KR", "LA1", "LA2", "NA1", "OC1", "PH2", "RU", "SG2", "TH2", "TR1", "TW2", "VN2"}
	CountryMap  = map[string]string{
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
)

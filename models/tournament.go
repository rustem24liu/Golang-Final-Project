package models

type Match struct {
	Team1  string `json:"team1"`
	Team2  string `json:"team2"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
}

type Round struct {
	Number  string  `json:"round_number"`
	Drawers []Match `json:"drawer"`
	Matches []Match `json:"matches"`
}

type TournamentResult struct {
	Winner string  `json:"winner"`
	Rounds []Round `json:"rounds"`
}

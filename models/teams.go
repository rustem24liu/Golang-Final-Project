package models

type Team struct {
	ID       int    `json:"team_id"`
	TeamName string `json:"team_name"`
	LeagueID int    `json:"league_id"`
}

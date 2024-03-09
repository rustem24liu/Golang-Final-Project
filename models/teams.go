package models

type Team struct {
	ID       int      `json:"id"`
	TeamName string   `json:"team_name"`
	Players  []Player `json:"players"`
}

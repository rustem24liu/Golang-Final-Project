package models

type Player struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	Cost      float64 `json:"cost"`
	Position  string  `json:"position"`
	TeamID    int     `json:"team_id"`
}

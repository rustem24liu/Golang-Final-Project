package models

type Coach struct {
	ID        int    `json:"coach_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ExpYear   int    `json:"exp_year"`
	TeamID    int    `json:"team_id"`
}

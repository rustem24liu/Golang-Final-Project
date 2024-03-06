package models

import "database/sql"

type Player struct {
	ID        int            `json:"id"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Age       int            `json:"age"`
	Cost      float64        `json:"cost"`
	Position  sql.NullString `json:"position"`
	TeamID    int            `json:"team_id"`
}

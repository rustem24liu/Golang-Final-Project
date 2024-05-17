package models

type Stadium struct {
	ID          int    `json:"id"`
	StadiumName string `json:"stadium_name"`
	Capacity    int    `json:"capacity"`
	TeamID      int    `json:"team_id"`
}

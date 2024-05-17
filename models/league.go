package models

type League struct {
	ID            int    `json:"id"`
	LeagueName    string `json:"league_name"`
	LeagueCountry string `json:"league_country"`
}

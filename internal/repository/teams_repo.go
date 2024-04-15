package repository

import (
	"database/sql"

	"github.com/rustem24liu/Golang-Final-Project/models"
)

type TeamRepo struct {
	db *sql.DB
}

func NewTeamRepo(db *sql.DB) *TeamRepo {
	return &TeamRepo{db}
}

func (r *TeamRepo) GetAllTeams() ([]models.Team, error) {
	query := `
        SELECT 
            t.team_id,
            t.team_name,
            p.player_id,
            p.first_name,
            p.last_name,
            p.player_age,
            p.player_cost,
            p.player_pos
        FROM 
            Teams t
        LEFT JOIN 
            Player p ON t.team_id = p.team_id
        ORDER BY 
            t.team_id, p.player_id
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teamsMap = make(map[int]models.Team)
	for rows.Next() {
		var teamID int
		var team models.Team
		var player models.Player
		err := rows.Scan(&teamID, &team.TeamName, &player.ID, &player.FirstName, &player.LastName, &player.Age, &player.Cost, &player.Position)
		if err != nil {
			return nil, err
		}

		if _, ok := teamsMap[teamID]; !ok {
			team.ID = teamID
			teamsMap[teamID] = team
		}

		updatedTeam := teamsMap[teamID]
		updatedTeam.Players = append(updatedTeam.Players, player)

		teamsMap[teamID] = updatedTeam
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var teams []models.Team
	for _, team := range teamsMap {
		teams = append(teams, team)
	}
	return teams, nil
}

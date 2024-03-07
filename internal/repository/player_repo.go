package repository

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	"github.com/rustem24liu/Golang-Final-Project/models"
	_ "github.com/rustem24liu/Golang-Final-Project/models"
)

type PlayerRepo struct {
	db *sql.DB
}

func NewPlayerRepo(db *sql.DB) *PlayerRepo {
	return &PlayerRepo{db}
}

func (r *PlayerRepo) GetAllPlayers() ([]models.Player, error) {
	rows, err := r.db.Query("SELECT * FROM Player")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Age, &player.Cost, &player.Position, &player.TeamID)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return players, nil
}

func (r *PlayerRepo) GetPlayerByID(id int) (*models.Player, error) {
	var player models.Player

	err := r.db.QueryRow("SELECT * FROM Player WHERE player_id = $1", id).
		Scan(&player.ID, &player.FirstName, &player.LastName, &player.Age, &player.Cost, &player.Position, &player.TeamID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("player not found")
		}
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepo) CreatePlayer(player *models.Player) error {
	_, err := r.db.Exec("INSERT INTO Player (first_name, last_name, player_age, player_cost, player_pos, team_id) VALUES ($1, $2, $3, $4, $5, $6)", player.FirstName, player.LastName, player.Age, player.Cost, player.Position, player.TeamID)
	if err != nil {
		fmt.Println("error")
	}
	return nil
}
func (r *PlayerRepo) UpdatePlayer(player models.Player) error {
	_, err := r.db.Exec("UPDATE Player SET first_name = $1, last_name = $2, player_age = $3, player_cost = $4, player_pos = $5, team_id = $6 WHERE player_id = $7", player.FirstName, player.LastName, player.Age, player.Cost, player.Position, player.TeamID, player.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *PlayerRepo) DeletePlayer(id int) error {
	_, err := r.db.Exec("Delete FROM Player WHERE player_id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil
}

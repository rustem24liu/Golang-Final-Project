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

func (r *PlayerRepo) ListOfAllPlayers() ([]models.Player, error) {
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

func (r *PlayerRepo) GetAllPlayers(pageNum, pageSize int, sortBy string, filters map[string]interface{}) ([]models.Player, error) {
	// Build SQL query based on sorting, filtering, and pagination parameters
	query := "SELECT * FROM Player"

	// Initialize slice to hold values for SQL query parameters
	var args []interface{}

	// Apply filtering if filters are provided
	if len(filters) > 0 {
		query += " WHERE "
		i := 1
		for key, value := range filters {
			// Type assertion to get the underlying value
			v, ok := value.(string)
			if !ok {
				// Handle the error if the assertion fails
				// For example, log an error or return an error response
				fmt.Printf("Error: Filter value for key %s is not a string\n", key)
				continue
			}
			// Use the value (v) as needed
			fmt.Printf("Applying filter: Key: %s, Value: %s\n", key, v)
			if i > 1 {
				query += " AND "
			}
			query += fmt.Sprintf("%s = $%d", key, i)
			fmt.Println(key)
			args = append(args, v)
			i++
		}
	}

	// Apply sorting
	if sortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", sortBy)
	}

	// Apply pagination
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, (pageNum-1)*pageSize)

	// Execute the query
	fmt.Println("Executing query:", query)
	fmt.Println("Query arguments:", args)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var players []models.Player
	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Age, &player.Cost, &player.Position, &player.TeamID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		players = append(players, player)
	}

	fmt.Println("Total players retrieved:", len(players))
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
	fmt.Println("Debugging: Inserting Player into database")
	fmt.Printf("Debugging: Player data - %+v\n", player)

	_, err := r.db.Exec("INSERT INTO Player (player_id, first_name, last_name, player_age, player_cost, player_pos, team_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", player.ID, player.FirstName, player.LastName, player.Age, player.Cost, player.Position, player.TeamID)
	if err != nil {
		fmt.Println("Error inserting player into database:", err)
		return err
	}

	fmt.Println("Debugging: Player inserted successfully")
	return nil
}

func (r *PlayerRepo) UpdatePlayer(player *models.Player) error {
	_, err := r.db.Exec("UPDATE Player SET first_name = $1, last_name = $2, player_age = $3, player_cost = $4, player_pos = $5, team_id = $6 WHERE player_id = $7", player.FirstName, player.LastName, player.Age, player.Cost, player.Position, player.TeamID, player.ID)
	if err != nil {
		return err
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

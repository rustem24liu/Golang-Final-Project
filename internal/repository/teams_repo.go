package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rustem24liu/Golang-Final-Project/models"
)

type TeamRepo struct {
	db *sql.DB
}

// NewTeamRepo creates a new TeamRepo instance.
func NewTeamRepo(db *sql.DB) *TeamRepo {
	return &TeamRepo{db}
}

func (r *TeamRepo) ListOfAllTeams() ([]models.Team, error) {
	rows, err := r.db.Query("SELECT * FROM Teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.ID, &team.TeamName, &team.LeagueID)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return teams, nil
}

func (r *TeamRepo) GetAllTeams(pageNum, pageSize int, sortBy string, filters map[string]interface{}) ([]models.Team, error) {

	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if sortBy == "" {
		sortBy = "team_id" // Default sort by ID
	}

	// Build SQL query based on sorting, filtering, and pagination parameters
	query := "SELECT * FROM Teams"
	fmt.Println("Query:", query)
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

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.ID, &team.TeamName, &team.LeagueID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		teams = append(teams, team)
	}

	fmt.Println("Total teams retrieved:", len(teams))
	return teams, nil
}

func (r *TeamRepo) GetTeamByID(id int) (*models.Team, error) {
	var team models.Team

	err := r.db.QueryRow("SELECT * FROM Teams WHERE team_id = $1", id).
		Scan(&team.ID, &team.TeamName, &team.LeagueID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("team not found")
		}
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepo) CreateTeam(team *models.Team) error {
	fmt.Println("Debugging: Inserting Player into database")
	fmt.Printf("Debugging: Player data - %+v\n", team)

	_, err := r.db.Exec("INSERT INTO Teams (team_name, league_id) VALUES ($1, $2)", team.TeamName, team.LeagueID)
	if err != nil {
		fmt.Println("Error inserting team into database:", err)
		return err
	}

	fmt.Println("Debugging: Team inserted successfully")
	return nil
}

func (r *TeamRepo) UpdateTeam(team *models.Team) error {
	_, err := r.db.Exec("UPDATE Teams SET team_name = $1, league_id = $2 WHERE team_id = $3", team.TeamName, team.LeagueID, team.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeamRepo) DeleteTeam(id int) error {
	_, err := r.db.Exec("DELETE FROM Teams WHERE team_id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil
}

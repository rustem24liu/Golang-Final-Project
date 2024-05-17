package repository

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"

	"github.com/rustem24liu/Golang-Final-Project/models"
	_ "github.com/rustem24liu/Golang-Final-Project/models"
)

type LeagueRepo struct {
	db *sql.DB
}

func NewLeagueRepo(db *sql.DB) *LeagueRepo {
	return &LeagueRepo{db}
}

func (r *LeagueRepo) ListOfAllLeagues() ([]models.League, error) {
	rows, err := r.db.Query("SELECT * FROM League")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leagues []models.League
	for rows.Next() {
		var league models.League
		err := rows.Scan(&league.ID, &league.LeagueName, &league.LeagueCountry)
		if err != nil {
			return nil, err
		}
		leagues = append(leagues, league)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return leagues, nil
}

func (r *LeagueRepo) GetAllLeagues(pageNum, pageSize int, sortBy string, filters map[string]interface{}) ([]models.League, error) {
	// Set default values if they are not provided
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if sortBy == "" {
		sortBy = "id" // Default sort by ID
	}

	// Build SQL query based on sorting, filtering, and pagination parameters
	query := "SELECT * FROM League"
	var args []interface{}

	// Apply filtering if filters are provided
	if len(filters) > 0 {
		query += " WHERE "
		i := 1
		for key, value := range filters {
			v, ok := value.(string)
			if !ok {
				fmt.Printf("Error: Filter value for key %s is not a string\n", key)
				continue
			}
			if i > 1 {
				query += " AND "
			}
			query += fmt.Sprintf("%s = $%d", key, i)
			args = append(args, v)
			i++
		}
	}

	// Apply sorting
	query += fmt.Sprintf(" ORDER BY %s", sortBy)

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

	var leagues []models.League
	for rows.Next() {
		var league models.League
		err := rows.Scan(&league.ID, &league.LeagueName, &league.LeagueCountry)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		leagues = append(leagues, league)
	}

	fmt.Println("Total leagues retrieved:", len(leagues))
	return leagues, nil
}

func (r *LeagueRepo) GetLeagueByID(id int) (*models.League, error) {
	var league models.League

	err := r.db.QueryRow("SELECT * FROM League WHERE id = $1", id).
		Scan(&league.ID, &league.LeagueName, &league.LeagueCountry)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("league not found")
		}
		return nil, err
	}
	return &league, nil
}

func (r *LeagueRepo) CreateLeague(league *models.League) error {
	fmt.Println("Debugging: Inserting League into database")
	fmt.Printf("Debugging: League data - %+v\n", league)

	_, err := r.db.Exec("INSERT INTO League (league_name, league_country) VALUES ($1, $2)", league.LeagueName, league.LeagueCountry)
	if err != nil {
		fmt.Println("Error inserting league into database:", err)
		return err
	}

	fmt.Println("Debugging: League inserted successfully")
	return nil
}

func (r *LeagueRepo) UpdateLeague(league *models.League) error {
	_, err := r.db.Exec("UPDATE League SET league_name = $1, league_country = $2 WHERE id = $3", league.LeagueName, league.LeagueCountry, league.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *LeagueRepo) DeleteLeague(id int) error {
	_, err := r.db.Exec("Delete FROM League WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil
}

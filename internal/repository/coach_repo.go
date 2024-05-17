package repository

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"

	"github.com/rustem24liu/Golang-Final-Project/models"
	_ "github.com/rustem24liu/Golang-Final-Project/models"
)

type CoachRepo struct {
	db *sql.DB
}

func NewCoachRepo(db *sql.DB) *CoachRepo {
	return &CoachRepo{db}
}

func (r *CoachRepo) ListOfAllCoaches() ([]models.Coach, error) {
	rows, err := r.db.Query("SELECT * FROM Coach")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coaches []models.Coach
	for rows.Next() {
		var coach models.Coach
		err := rows.Scan(&coach.ID, &coach.FirstName, &coach.LastName, &coach.ExpYear, &coach.TeamID)
		if err != nil {
			return nil, err
		}
		coaches = append(coaches, coach)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return coaches, nil
}

func (r *CoachRepo) GetAllCoaches(pageNum, pageSize int, sortBy string, filters map[string]interface{}) ([]models.Coach, error) {

	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if sortBy == "" {
		sortBy = "coach_id"
	}

	// Build SQL query based on sorting, filtering, and pagination parameters
	query := "SELECT * FROM Coach"
	fmt.Printf("Query is the", query)
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

	var coaches []models.Coach
	for rows.Next() {
		var coach models.Coach
		err := rows.Scan(&coach.ID, &coach.FirstName, &coach.LastName, &coach.ExpYear, &coach.TeamID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		coaches = append(coaches, coach)
	}

	fmt.Println("Total coaches retrieved:", len(coaches))
	return coaches, nil
}

func (r *CoachRepo) GetCoachByID(id int) (*models.Coach, error) {
	var coach models.Coach

	err := r.db.QueryRow("SELECT * FROM Coach WHERE coach_id = $1", id).
		Scan(coach.ID, &coach.FirstName, &coach.LastName, &coach.ExpYear, &coach.TeamID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("coach not found")
		}
		return nil, err
	}
	return &coach, nil
}

func (r *CoachRepo) CreateCoach(coach *models.Coach) error {
	fmt.Println("Debugging: Inserting Coach into database")
	fmt.Printf("Debugging: Coach data - %+v\n", coach)

	_, err := r.db.Exec("INSERT INTO Coach (first_name, last_name, exp_year, team_id) VALUES ($1, $2, $3, $4)", coach.FirstName, coach.LastName, coach.ExpYear, coach.TeamID)
	if err != nil {
		fmt.Println("Error inserting coach into database:", err)
		return err
	}

	fmt.Println("Debugging: League inserted successfully")
	return nil
}

func (r *CoachRepo) UpdateCoach(coach *models.Coach) error {
	_, err := r.db.Exec("UPDATE Coach SET first_name = $1, last_name = $2, exp_year = $3 , team_id = $4 WHERE league_id = $5", coach.FirstName, coach.LastName, coach.ExpYear, coach.TeamID, coach.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CoachRepo) DeleteCoach(id int) error {
	_, err := r.db.Exec("DELETE FROM Coach WHERE coach_id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil
}

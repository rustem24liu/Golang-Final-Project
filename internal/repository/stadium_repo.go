package repository

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"

	"github.com/rustem24liu/Golang-Final-Project/models"
	_ "github.com/rustem24liu/Golang-Final-Project/models"
)

type StadiumRepo struct {
	db *sql.DB
}

func NewStadiumRepo(db *sql.DB) *StadiumRepo {
	return &StadiumRepo{db}
}

func (r *StadiumRepo) ListOfAllStadiums() ([]models.Stadium, error) {
	rows, err := r.db.Query("SELECT * FROM Stadiums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stadiums []models.Stadium
	for rows.Next() {
		var stadium models.Stadium
		err := rows.Scan(&stadium.ID, &stadium.StadiumName, &stadium.Capacity, &stadium.TeamID)
		if err != nil {
			return nil, err
		}
		stadiums = append(stadiums, stadium)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return stadiums, nil
}

func (r *StadiumRepo) GetAllStadiums(pageNum, pageSize int, sortBy string, filters map[string]interface{}) ([]models.Stadium, error) {

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
	query := "SELECT * FROM Stadiums"
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

	var stadiums []models.Stadium
	for rows.Next() {
		var stadium models.Stadium
		err := rows.Scan(&stadium.ID, &stadium.StadiumName, &stadium.Capacity, &stadium.TeamID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		stadiums = append(stadiums, stadium)
	}

	fmt.Println("Total stadiums retrieved:", len(stadiums))
	return stadiums, nil
}

func (r *StadiumRepo) GetStadiumByID(id int) (*models.Stadium, error) {
	var stadium models.Stadium

	err := r.db.QueryRow("SELECT * FROM Stadiums WHERE id = $1", id).
		Scan(&stadium.ID, &stadium.StadiumName, &stadium.Capacity, &stadium.TeamID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("stadium not found")
		}
		return nil, err
	}
	return &stadium, nil
}

func (r *StadiumRepo) CreateStadium(stadium *models.Stadium) error {
	fmt.Println("Debugging: Inserting Stadium into database")
	fmt.Printf("Debugging: Stadium data - %+v\n", stadium)

	_, err := r.db.Exec("INSERT INTO Stadiums (stadium_name, capacity, team_id) VALUES ($1, $2, $3)", stadium.StadiumName, stadium.Capacity, stadium.TeamID)
	if err != nil {
		fmt.Println("Error inserting stadium into database:", err)
		return err
	}

	fmt.Println("Debugging: Stadium inserted successfully")
	return nil
}

func (r *StadiumRepo) UpdateStadium(stadium *models.Stadium) error {
	_, err := r.db.Exec("UPDATE Stadiums SET stadium_name = $1, capacity = $2, team_id = $3 WHERE league_id = $4", stadium.StadiumName, stadium.Capacity, stadium.TeamID, stadium.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *StadiumRepo) DeleteStadium(id int) error {
	_, err := r.db.Exec("DELETE FROM Stadiums WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return nil
}

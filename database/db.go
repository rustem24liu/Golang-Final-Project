package database

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "0510"
	dbname   = "football_team"
)

func GetTeamNames() ([]string, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT team_name FROM teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []string
	for rows.Next() {
		var team_name string
		if err := rows.Scan(&team_name); err != nil {
			return nil, err
		}
		teams = append(teams, team_name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return teams, nil
}

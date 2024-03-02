package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type Match struct {
	Team1  string
	Team2  string
	Score1 int
	Score2 int
}

func main() {
<<<<<<< HEAD:cmd/main.go
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://postgres:ayan2004@localhost/Go-24?sslmode=disable")
=======
	// Establish a connection to the PostgreSQL database
	db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
>>>>>>> 086a62f4289f3dba004a98603a4f110cc29d89dc:main.go
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Fetch team names from the database
	rows, err := db.Query("SELECT team_name FROM teams")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var teams []string
	for rows.Next() {
		var teamName string
<<<<<<< HEAD:cmd/main.go
		if err := rows.Scan(&teamName); err != nil {
=======
		err := rows.Scan(&teamName)
		if err != nil {
>>>>>>> 086a62f4289f3dba004a98603a4f110cc29d89dc:main.go
			panic(err)
		}
		teams = append(teams, teamName)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})

	var winners []string
<<<<<<< HEAD:cmd/main.go
	var drawers []Match // Save matches that ended in a draw

	for len(teams) > 1 {
		var nextRound []string

=======
	var drawers []Match
	comments := []string{"Round of 16", "Quarterfinal", "Halffinal", "Final"}
	cnt := 0

	for len(teams) > 1 {
		var nextRound []string
		var matches []Match

		fmt.Println(comments[cnt])
>>>>>>> 086a62f4289f3dba004a98603a4f110cc29d89dc:main.go
		for i := 0; i < len(teams); i += 2 {
			match := Match{Team1: teams[i], Team2: teams[i+1]}
			match.Score1 = rand.Intn(5)
			match.Score2 = rand.Intn(5)
			fmt.Printf("%s %d - %d %s\n", match.Team1, match.Score1, match.Score2, match.Team2)

			if match.Score1 == match.Score2 {
				drawers = append(drawers, match)
			} else if match.Score1 > match.Score2 {
				winners = append(winners, match.Team1)
				nextRound = append(nextRound, match.Team1)
			} else {
				winners = append(winners, match.Team2)
				nextRound = append(nextRound, match.Team2)
			}
		}

		for _, drawMatch := range drawers {
			for {
				drawMatch.Score1 = rand.Intn(5)
				drawMatch.Score2 = rand.Intn(5)
				fmt.Printf("%s %d - %d %s\n", drawMatch.Team1, drawMatch.Score1, drawMatch.Score2, drawMatch.Team2)

				if drawMatch.Score1 != drawMatch.Score2 {
					if drawMatch.Score1 > drawMatch.Score2 {
						winners = append(winners, drawMatch.Team1)
						nextRound = append(nextRound, drawMatch.Team1)
					} else {
						winners = append(winners, drawMatch.Team2)
						nextRound = append(nextRound, drawMatch.Team2)
					}
					break
				}
			}
		}

		teams = nextRound
		drawers = nil
<<<<<<< HEAD:cmd/main.go
		fmt.Println("Next Round:")
	}

=======
		cnt++
	}

	// Print the winner
>>>>>>> 086a62f4289f3dba004a98603a4f110cc29d89dc:main.go
	fmt.Println("Winner:")
	fmt.Println(teams[0])
}

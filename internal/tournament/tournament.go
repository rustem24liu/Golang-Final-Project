package tournament

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
)

type Match struct {
	Team1  string
	Team2  string
	Score1 int
	Score2 int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "football_team"
)

func RunTournament() {
	rand.Seed(time.Now().UnixNano())

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT team_name FROM teams")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var teams []string

	for rows.Next() {
		var team_name string
		if err := rows.Scan(&team_name); err != nil {
			panic(err)
		}
		teams = append(teams, team_name)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for len(teams) > 1 {
		rand.Shuffle(len(teams), func(i, j int) {
			teams[i], teams[j] = teams[j], teams[i]
		})

		var nextRound []string
		var matches []Match

		fmt.Println("Next Round:")

		for i := 0; i < len(teams); i += 2 {
			if i+1 < len(teams) {
				match := Match{Team1: teams[i], Team2: teams[i+1]}
				match.Score1 = rand.Intn(5)
				match.Score2 = rand.Intn(5)
				fmt.Printf("%s %d - %d %s\n", match.Team1, match.Score1, match.Score2, match.Team2)

				if match.Score1 == match.Score2 {
					// Draw
				} else if match.Score1 > match.Score2 {
					nextRound = append(nextRound, match.Team1)
				} else {
					nextRound = append(nextRound, match.Team2)
				}

				matches = append(matches, match)
			} else {
				// Handle odd number of teams (bye or advancing automatically)
				nextRound = append(nextRound, teams[i])
			}
		}

		teams = nextRound

		if len(teams) == 1 {
			break
		}
	}

	if len(teams) > 0 {
		fmt.Println("Winner:")
		fmt.Println(teams[0])
	} else {
		fmt.Println("No teams remaining.")
	}
}

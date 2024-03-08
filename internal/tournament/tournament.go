package tournament

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

type Match struct {
	Team1  string `json:"team1"`
	Team2  string `json:"team2"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
}

type TournamentResult struct {
	Winner  string  `json:"winner"`
	Matches []Match `json:"matches"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1000tenge"
	dbname   = "football_team"
)

func RunTournament() TournamentResult {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
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
		if err := rows.Scan(&teamName); err != nil {
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

	var tournamentResult TournamentResult
	var winners []string
	var drawers []Match
	cnt := 0

	var matches []Match

	for len(teams) > 1 {
		var nextRound []string

		for i := 0; i < len(teams); i += 2 {
			match := Match{Team1: teams[i], Team2: teams[i+1]}
			match.Score1 = rand.Intn(5)
			match.Score2 = rand.Intn(5)

			// Determine winners and next round participants
			if match.Score1 == match.Score2 {
				drawers = append(drawers, match)
			} else if match.Score1 > match.Score2 {
				winners = append(winners, match.Team1)
				nextRound = append(nextRound, match.Team1)
			} else {
				winners = append(winners, match.Team2)
				nextRound = append(nextRound, match.Team2)
			}

			matches = append(matches, match) // Append match to matches slice
		}
		for _, drawMatch := range drawers {
			for {
				drawMatch.Score1 = rand.Intn(5)
				drawMatch.Score2 = rand.Intn(5)

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
		cnt++
	}

	tournamentResult.Winner = teams[0]
	tournamentResult.Matches = matches

	return tournamentResult
}

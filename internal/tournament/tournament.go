package tournament

import (
	"database/sql"
	"fmt"
	"github.com/rustem24liu/Golang-Final-Project/models"
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/rustem24liu/Golang-Final-Project/models"
)

func RunTournament() (models.TournamentResult, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return models.TournamentResult{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT team_name FROM teams")
	if err != nil {
		return models.TournamentResult{}, err
	}
	defer rows.Close()

	var teams []string
	for rows.Next() {
		var teamName string
		if err := rows.Scan(&teamName); err != nil {
			return models.TournamentResult{}, err
		}
		teams = append(teams, teamName)
	}

	if err := rows.Err(); err != nil {
		return models.TournamentResult{}, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})

	var tournamentResult models.TournamentResult
	comments := []string{"Round of 8", "Quarter final", "Semi final", "Final"}
	var winners []string
	var drawers []models.Match
	var drawersResult []models.Match

	var rounds []models.Round

	for len(teams) > 1 {
		var nextRound []string
		var matches []models.Match

		for i := 0; i < len(teams); i += 2 {
			match := models.Match{Team1: teams[i], Team2: teams[i+1]}
			match.Score1 = rand.Intn(5)
			match.Score2 = rand.Intn(5)

			if match.Score1 == match.Score2 {
				drawers = append(drawers, match)
			} else if match.Score1 > match.Score2 {
				winners = append(winners, match.Team1)
				nextRound = append(nextRound, match.Team1)
			} else {
				winners = append(winners, match.Team2)
				nextRound = append(nextRound, match.Team2)
			}

			matches = append(matches, match)
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
					drawersResult = append(drawersResult, drawMatch)
					break
				}
			}
		}
		teams = nextRound
		drawers = nil

		round := models.Round{
			Number:  comments[len(rounds)],
			Drawers: drawersResult,
			Matches: matches,
		}
		rounds = append(rounds, round)
	}

	tournamentResult.Winner = teams[0]
	tournamentResult.Rounds = rounds

	return tournamentResult, nil
}

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
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://postgres:ayan2004@localhost/Go-24?sslmode=disable")
	// Establish a connection to the PostgreSQL database
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

		err := rows.Scan(&teamName)
		if err != nil {
			if err := rows.Scan(&teamName); err != nil {
				err := rows.Scan(&teamName)
				if err != nil {
					panic(err)
				}
				teams = append(teams, teamName)
			}
			if err := rows.Err(); err != nil {
				panic(err)
			}

			// Shuffle the teams
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(teams), func(i, j int) {
				teams[i], teams[j] = teams[j], teams[i]
			})

			var winners []string
			var drawers []Match // Save matches that ended in a draw

			// Simulate rounds until there's only one team left
			for len(teams) > 1 {
				var nextRound []string
				var matches []Match

				// Generate matches for the current round

				var drawers []Match
				comments := []string{"Round of 16", "Quarterfinal", "Halffinal", "Final"}
				cnt := 0

				for len(teams) > 1 {
					var nextRound []string
					var matches []Match

					fmt.Println(comments[cnt])
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

						matches = append(matches, match)
					}

					// Handle matches that ended in a draw
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

					// Prepare for the next round
					teams = nextRound
					drawers = nil // Reset drawers for the next round
					fmt.Println("Next Round:")
				}

				// Print the winner

				drawers = nil
				fmt.Println("Next Round:")
				cnt++
			}

		}

		// Print the winner
		fmt.Println("Winner:")
		fmt.Println(teams[0])
	}
}

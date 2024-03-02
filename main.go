package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

type Match struct {
	Team1  string
	Team2  string
	Score1 int
	Score2 int
}

func main() {
	db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
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

	var winners []string
	var drawers []Match // Save matches that ended in a draw

<<<<<<< HEAD
<<<<<<< HEAD
		for _, drawMatch := range drawers {
			for {
				drawMatch.Score1 = rand.Intn(5)
				drawMatch.Score2 = rand.Intn(5)
				fmt.Printf("%s %d - %d %s\n", drawMatch.Team1, drawMatch.Score1, drawMatch.Score2, drawMatch.Team2)
=======
    // Simulate rounds until there's only one team left
    for len(teams) > 1 {
        var nextRound []string
        var matches []Match
>>>>>>> fca792782f4c7e1b1d32bae918a8e34a8344b5f9
=======
	for len(teams) > 1 {
		var nextRound []string
>>>>>>> 638440529321f0baba2361864384b9af88dbf9cd

		for i := 0; i < len(teams); i += 2 {
			match := Match{Team1: teams[i], Team2: teams[i+1]}
			match.Score1 = rand.Intn(5)
			match.Score2 = rand.Intn(5)
			fmt.Printf("%s %d - %d %s\n", match.Team1, match.Score1, match.Score2, match.Team2)

<<<<<<< HEAD
<<<<<<< HEAD
		teams = nextRound
		drawers = nil
		cnt++
	}
=======
            if match.Score1 == match.Score2 {
                drawers = append(drawers, match)
            } else if match.Score1 > match.Score2 {
                winners = append(winners, match.Team1)
                nextRound = append(nextRound, match.Team1)
            } else {
                winners = append(winners, match.Team2)
                nextRound = append(nextRound, match.Team2)
            }
>>>>>>> fca792782f4c7e1b1d32bae918a8e34a8344b5f9
=======
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
>>>>>>> 638440529321f0baba2361864384b9af88dbf9cd

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
		fmt.Println("Next Round:")
	}

	fmt.Println("Winner:")
	fmt.Println(teams[0])
}
}
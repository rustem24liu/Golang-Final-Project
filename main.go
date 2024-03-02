package main

import (
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
	teams := []string{"Team 1", "Team 2", "Team 3", "Team 4", "Team 5", "Team 6", "Team 7", "Team 8", "Team 9", "Team 10", "Team 11", "Team 12", "Team 13", "Team 14", "Team 15", "Team 16"}
	comments := []string{"Round of 16", "Quarterfinal", "Halffinal", "Final"}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})

	var winners []string
	var drawers []Match
	cnt := 0
	// Simulate rounds until there's only one team left
	for len(teams) > 1 {
		var nextRound []string
		var matches []Match
		fmt.Println(comments[cnt])
		// Generate matches for the current round
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
		cnt++
	}

	// Print the winner
	fmt.Println()
	fmt.Println("Winner:")
	fmt.Println(teams[0])
}

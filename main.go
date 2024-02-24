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
	teams := []string{"Team 1", "Team 2", "Team 3", "Team 4", "Team 5", "Team 6", "Team 7", "Team 8"}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})

	quarterFinals := make([]Match, 0)
	for i := 0; i < len(teams); i += 2 {
		match := Match{Team1: teams[i], Team2: teams[i+1]}
		quarterFinals = append(quarterFinals, match)
	}

	var winners []string
	var drawers []Match // Сохраняем матчи, которые закончились вничью

	fmt.Println("Quarterfinals:")
	for _, match := range quarterFinals {
		match.Score1 = rand.Intn(5)
		match.Score2 = rand.Intn(5)
		fmt.Printf("%s %d - %d %s\n", match.Team1, match.Score1, match.Score2, match.Team2)

		if match.Score1 == match.Score2 {
			drawers = append(drawers, match)
		} else if match.Score1 > match.Score2 {
			winners = append(winners, match.Team1)
		} else {
			winners = append(winners, match.Team2)
		}
	}

	fmt.Println("Rematches for draws:")
	for _, drawMatch := range drawers {
		for {
			drawMatch.Score1 = rand.Intn(5)
			drawMatch.Score2 = rand.Intn(5)
			fmt.Printf("%s %d - %d %s\n", drawMatch.Team1, drawMatch.Score1, drawMatch.Score2, drawMatch.Team2)

			if drawMatch.Score1 != drawMatch.Score2 {
				if drawMatch.Score1 > drawMatch.Score2 {
					winners = append(winners, drawMatch.Team1)
				} else {
					winners = append(winners, drawMatch.Team2)
				}
				break
			}
		}
	}

	fmt.Println("Winners of the quarterfinals:")
	for _, winner := range winners {
		fmt.Println(winner)
	}
}

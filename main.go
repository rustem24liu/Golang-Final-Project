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

	for i := range quarterFinals {
		quarterFinals[i].Score1 = rand.Intn(5)
		quarterFinals[i].Score2 = rand.Intn(5)
	}

	fmt.Println("Quarterfinals:")
	for _, match := range quarterFinals {
		fmt.Printf("%s %d - %d %s\n", match.Team1, match.Score1, match.Score2, match.Team2)
	}
}

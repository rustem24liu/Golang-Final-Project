package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/lib/pq"
	"github.com/rustem24liu/Golang-Final-Project/internal/tournament"
	_ "github.com/rustem24liu/Golang-Final-Project/internal/tournament"
)

func main() {
	tournament.RunTournament()
	fmt.Println("Tournament finished")
}

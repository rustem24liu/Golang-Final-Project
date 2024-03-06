package main

import (
	"database/sql"
	_ "database/sql"
	_ "encoding/json"
	_ "github.com/rustem24liu/Golang-Final-Project/internal/repository"
	_ "log"
	_ "net/http"
)

type PlayerHandler struct {
	playerRepo *PlayerRepo
}

func NewPlayerHandler(db *sql.DB) *PlayerHandler {
	return &PlayerHandler{
		playerRepo: NewPlayerRepo(db),
	}
}

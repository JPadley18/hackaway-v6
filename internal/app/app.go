package app

import "cordle/internal/database"

const (
	dbKey = "conf/db-key.json"
	discTok = "conf/discord-tok.json"
)

func Run() {
	db := database.NewDb(dbKey)
	defer db.Close()

	
}
package app

import "cordle/internal/database"

const dbKey = "conf/db-key.json"

func Run() {
	d := database.NewDb(dbKey)
	defer d.Close()
}
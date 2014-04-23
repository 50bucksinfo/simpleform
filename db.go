package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func connectToDB() {
	var err error
	db, err = sql.Open("postgres", config.ConnectionString)
	logFatal(err, "database connect error")
	logInfo("Connected to the database using ", config.ConnectionString)
}

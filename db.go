package main

import (
	"database/sql"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

var db *sql.DB

func connectToDB() {
	var err error
	db, err = sql.Open("postgres", config.ConnectionString)
	if err != nil {
		glog.Fatalln(err, "database connect error")
	}
	glog.Infoln("Connected to the database using ", config.ConnectionString)
}

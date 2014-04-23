package main

import (
	"net/http"
	"time"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	//validate email
	//create form api token and api token
	//create user in db
	email := r.FormValue("email")
	var id int
	err := db.QueryRow("INSERT INTO users(email, form_api_token, api_token, created_at, updated_at) VALUES($1, $2, $3, $4, $4) RETURNING ID",
		email, secureHex(), secureHex(), time.Now().UTC()).Scan(&id)

	logInfo("created user with id", id)

	logError(err)

	//err = db.QueryRow("INSERT INTO forms(site_id, entry, request_ip, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING ID",
	//siteID, string(data), r.RemoteAddr, formName, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	//send user a notification via email
	//show him a page on how to use this new token
	w.Write([]byte("awesome"))
}

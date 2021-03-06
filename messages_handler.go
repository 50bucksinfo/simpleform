package main

import (
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type message struct {
	ID                                                int
	FormApiToken, Data, RequestIP, Referrer, FormName string
	FormData                                          url.Values
	Spam                                              bool
	CreatedAt                                         time.Time
}

type messageJson struct {
	ID                            int
	RequestIP, Referrer, FormName string
	FormData                      map[string]string
	Spam                          bool
	CreatedAt                     time.Time
}

//TODO add pagination
func messagesIndexHandler(w http.ResponseWriter, r *http.Request) {
	apiToken := r.FormValue("api_token")
	messages, err := getMessages(apiToken)

	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occured"))
		return
	}

	render(w, "messages.html", messages)
}

func messagesIndexJsonHandler(w http.ResponseWriter, r *http.Request) {
	apiToken := r.FormValue("api_token")
	messages, err := getMessages(apiToken)

	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occured"))
		return
	}

	jsonMessages := make([]messageJson, 0, len(messages))
	for _, message := range messages {
		mj := messageJson{ID: message.ID, RequestIP: message.RequestIP, Referrer: message.Referrer, FormName: message.FormName, Spam: message.Spam, CreatedAt: message.CreatedAt}
		mj.FormData = make(map[string]string)
		for k, v := range message.FormData {
			mj.FormData[k] = strings.Join(v, ",")
		}
		jsonMessages = append(jsonMessages, mj)
	}

	messageJson, err := json.Marshal(jsonMessages)

	if err != nil {
		logError(err, "JSON MARSHAL ERROR")
	}

	w.Write(messageJson)
}

func getMessages(apiToken string) ([]message, error) {
	messages := make([]message, 0, 100)
	rows, err := db.Query("SELECT id, form_api_token, data, request_ip, referrer, form_name, created_at  FROM messages WHERE form_api_token = (SELECT form_api_token FROM users WHERE api_token = $1 LIMIT 1)", apiToken)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		m := &message{}
		err = rows.Scan(&m.ID, &m.FormApiToken, &m.Data, &m.RequestIP, &m.Referrer, &m.FormName, &m.CreatedAt)
		if err != nil {
			logError(err, "messages index row.scan")
			continue
		}
		//TODO move this to a message method
		m.FormData = make(url.Values)
		json.Unmarshal([]byte(m.Data), &m.FormData)
		messages = append(messages, *m)
	}
	return messages, nil
}

//form post handler start
func messageHandler(w http.ResponseWriter, r *http.Request) {
	//parse the form
	r.ParseMultipartForm(32 << 20) //32 MB

	formName := withDefault(r.FormValue("_name"), r.URL.Path[1:], "Default Form")
	formName = normalizeName(formName)
	r.Form.Del("_name")

	redirect := withDefault(r.Referer(), "http://"+r.Host+r.URL.Path)
	redirect += "#thank-you"
	redirect = withDefault(r.FormValue("_redirect"), r.FormValue("redirect_to "), redirect)
	r.Form.Del("_redirect")
	r.Form.Del("redirect_to")

	formApiToken := r.FormValue("form_api_token")
	r.Form.Del("form_api_token")

	data, err := json.Marshal(r.Form)
	logError(err)

	//validate api token
	var dummy int
	err = db.QueryRow("SELECT 1 FROM users WHERE form_api_token = $1", formApiToken).Scan(&dummy)

	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid api token"))
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}

	var id int
	err = db.QueryRow("INSERT INTO messages(form_api_token, data, request_ip, referrer, form_name, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING ID",
		formApiToken, string(data), ip, r.Referer(), formName, time.Now().UTC()).Scan(&id)
	logInfo("inserted form", id)

	//handle spam prevention
	//postProcessForm(id)

	logError(err, "ERROR in INSERT")

	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Redirect(w, r, redirect, http.StatusMovedPermanently)
	}
}

var formNameWhiteList = regexp.MustCompile("[^a-zA-Z0-9]+")

func normalizeName(name string) string {
	name = formNameWhiteList.ReplaceAllString(name, " ")
	return strings.Trim(name, " ")
}

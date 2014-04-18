package main

import (
	"encoding/json"
	"github.com/golang/glog"
	"net/http"
	"regexp"
	"strings"
	"time"
)

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
	if err != nil {
		glog.Errorln(err)
	}

	//validate api token
	var dummy int
	err = db.QueryRow("SELECT 1 FROM users WHERE form_api_token = $1", formApiToken).Scan(&dummy)

	if err != nil {
		glog.Errorln(err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid api token"))
		return
	}

	var id int
	err = db.QueryRow("INSERT INTO messages(form_api_token, data, request_ip, referrer, form_name, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING ID",
		formApiToken, string(data), r.RemoteAddr, r.Referer(), formName, time.Now().UTC()).Scan(&id)
	glog.Infoln("inserted form", id)

	//handle spam prevention
	//postProcessForm(id)

	if err != nil {
		glog.Errorln(err, "ERROR in INSERT")
	}

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

//form post handler end
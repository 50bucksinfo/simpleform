package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
	"time"
)

var configPath string

func main() {
	defer glog.Flush()
	flag.StringVar(&configPath, "config", "./config.json", "config file path")
	flag.Parse()
	loadConfig()
	loadTemplates()
	connectToDB()
	startServer()
}

//http server
func startServer() {
	glog.Infof("Starting server at http://localhost%s/\n", config.Port)
	r := mux.NewRouter()
	wireupRoutes(r)
	glog.Fatalln(http.ListenAndServe(config.Port, r), "failed to start server")
}

//TODO do we really need mux here, if not remove it
func wireupRoutes(r *mux.Router) {
	//GET
	r.HandleFunc("/", viewHandler("index.html")).Methods("GET")
	r.HandleFunc("/about", viewHandler("about.html")).Methods("GET")
	r.HandleFunc("/demo", viewHandler("demo.html")).Methods("GET")

	//POST
	r.HandleFunc("/users", createUserHandler).Methods("POST")
}

func viewHandler(view string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, view, "")
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infoln("create user handler")
	//validate email
	//create form api token and api token
	//create user in db
	email := r.FormValue("email")
	var id int
	err := db.QueryRow("INSERT INTO users(email, form_api_token, api_token, created_at, updated_at) VALUES($1, $2, $3, $4, $4) RETURNING ID",
		email, secureHex(), secureHex(), time.Now().UTC()).Scan(&id)

	glog.Infoln("created user with id", id)

	if err != nil {
		glog.Infoln(err)
	}

	//err = db.QueryRow("INSERT INTO forms(site_id, entry, request_ip, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING ID",
	//siteID, string(data), r.RemoteAddr, formName, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	//send user a notification via email
	//show him a page on how to use this new token
	w.Write([]byte("awesome"))
}

//templates
var templates *template.Template

func render(w io.Writer, view string, data interface{}) {
	buf := bytes.NewBufferString("")
	err := templates.ExecuteTemplate(buf, view, data)
	if err != nil {
		glog.Errorln(err, "failed to render view:", view)
	}
	err = templates.ExecuteTemplate(w, "layout.html", template.HTML(buf.String()))
	if err != nil {
		glog.Errorln(err, "failed to render view: layout.html")
	}
}

func loadTemplates() {
	var err error
	//TODO use os.Home
	templates, err = template.ParseGlob("/home/minhajuddin/gocode/src/github.com/minhajuddin/simpleform/views/*.html")
	if err != nil {
		glog.Fatalln(err, "failed to load templates")
	}
}

//random
func secureHex() string {
	hexBytes := make([]byte, 32)
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		glog.Errorln(err)
		return ""
	}
	hex.Encode(hexBytes, randomBytes)
	return string(hexBytes)
}

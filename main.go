package main

import (
	"bytes"
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
)

var configPath string

func main() {
	defer glog.Flush()
	flag.StringVar(&configPath, "config", "./config.json", "config file path")
	flag.Parse()
	loadConfig()
	loadTemplates()
	startServer()
}

//http server
func startServer() {
	glog.Infof("Starting server at http://localhost%s/\n", config.Port)
	r := mux.NewRouter()
	wireupRoutes(r)
	glog.Fatalln(http.ListenAndServe(config.Port, r), "failed to start server")
}

func wireupRoutes(r *mux.Router) {
	r.HandleFunc("/", viewHandler("index.html"))
	r.HandleFunc("/about", viewHandler("about.html"))
}

func viewHandler(view string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, view, "")
	}
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

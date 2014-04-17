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

func main() {
	defer glog.Flush()
	flag.Parse()
	loadTemplates()
	startServer()
}

func startServer() {
	glog.Infoln("Starting server at http://localhost:3030/")
	r := mux.NewRouter()
	wireupRoutes(r)
	glog.Fatalln(http.ListenAndServe(":3030", r), "failed to start server")
}

func wireupRoutes(r *mux.Router) {
	r.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", "")
}

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
	glog.Infoln(templates.Name(), "templates")
}

package main

import (
	"bytes"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
)

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
	r.HandleFunc("/messages", messageHandler).Methods("POST")
}

func viewHandler(view string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, view, config)
	}
}

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

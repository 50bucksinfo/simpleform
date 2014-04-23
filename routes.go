package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
)

//http server
func startServer() {
	logInfo("Starting server at http://localhost%s/\n", config.Port)
	r := mux.NewRouter()
	wireupRoutes(r)
	logFatal(http.ListenAndServe(config.Port, r), "failed to start server")
}

//TODO do we really need mux here, if not remove it
func wireupRoutes(r *mux.Router) {
	//GET
	r.HandleFunc("/", viewHandler("index.html")).Methods("GET")
	r.HandleFunc("/about", viewHandler("about.html")).Methods("GET")
	r.HandleFunc("/demo", viewHandler("demo.html")).Methods("GET")
	r.HandleFunc("/messages", messagesIndexHandler).Methods("GET")
	r.HandleFunc("/messages.json", messagesIndexJsonHandler).Methods("GET")

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
	logError(err, "failed to render view:", view)
	err = templates.ExecuteTemplate(w, "layout.html", template.HTML(buf.String()))
	logError(err, "failed to render view: layout.html")
}

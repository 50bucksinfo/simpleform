package main

import (
	"html/template"
	"strings"
)

var templates *template.Template

func loadTemplates() {
	var err error
	fns := template.FuncMap{
		"join": strings.Join,
	}
	templates, err = template.New("").Funcs(fns).ParseGlob("./views/*.html")
	logFatal(err, "failed to load templates")
}

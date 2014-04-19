package main

import (
	"github.com/golang/glog"
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
	if err != nil {
		glog.Fatalln(err, "failed to load templates")
	}
}

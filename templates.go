package main

import (
	"github.com/golang/glog"
	"html/template"
	"strings"
)

var templates *template.Template

func loadTemplates() {
	var err error
	templates, err = template.ParseGlob("./views/*.html")
	if err != nil {
		glog.Fatalln(err, "failed to load templates")
	}
	templates = templates.Funcs(fns)
}

var fns = template.FuncMap{
	"join":  func() string { return "HEE" },
	"title": strings.Title,
}

package main

import (
	"github.com/golang/glog"
	"html/template"
)

var templates *template.Template

func loadTemplates() {
	var err error
	templates, err = template.ParseGlob("./views/*.html")
	if err != nil {
		glog.Fatalln(err, "failed to load templates")
	}
}

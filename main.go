package main

import (
	"flag"
	"github.com/golang/glog"
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

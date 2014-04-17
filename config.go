package main

import (
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
)

var config struct {
	Port             string `json:"port"`
	ConnectionString string `json:"connectionString"`
	Host             string `json:"host"`
}

func loadConfig() {
	configBytes, err := ioutil.ReadFile("./config.json")
	glog.Infoln("CONFIG: ", string(configBytes))
	if err != nil {
		glog.Fatalln(err, "config file not found")
	}
	json.Unmarshal(configBytes, &config)
	glog.Infoln("loaded config: ", config)
}

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
	Demo             struct {
		Email        string `json:"email"`
		ApiToken     string `json:"apiToken"`
		FormApiToken string `json:"formApiToken"`
	} `json:"demo"`
}

func loadConfig() {
	configBytes, err := ioutil.ReadFile("./config.json")
	glog.Infoln("CONFIG: ", string(configBytes))
	if err != nil {
		glog.Fatalln(err, "config file not found")
	}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		glog.Fatalln(err, "json parse error")
	}
	glog.Infoln("loaded config: ", config)
}

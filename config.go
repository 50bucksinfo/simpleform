package main

import (
	"encoding/json"
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
	configBytes, err := ioutil.ReadFile(configPath)
	logInfo("CONFIG: ", string(configBytes))
	logFatal(err, "config file not found")

	err = json.Unmarshal(configBytes, &config)
	logFatal(err, "json parse error")

	logInfo("loaded config: ", config)
}

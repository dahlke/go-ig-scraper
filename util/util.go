package util

import (
	"encoding/json"
	"io/ioutil"
	"os"

	structs "github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

func ConfigLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func ParseConfig(configJSONPath string) structs.AppConfig {
	configJSONFile, err := os.Open(configJSONPath)
	if err != nil {
		log.Error(err)
	}
	defer configJSONFile.Close()

	jsonBytes, _ := ioutil.ReadAll(configJSONFile)
	var config structs.AppConfig
	json.Unmarshal(jsonBytes, &config)

	return config
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var ConfigVar = getConfig()

func getConfig() Config {

	var config Config

	configFile, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalf("Failed to open ConfigVar file: %v", err)
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	err = yaml.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal ConfigVar: %v", err)
	}

	return config
}

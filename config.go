package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var config = getConfig()

func getConfig() Config {
	var config Config

	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	err = yaml.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return config
}

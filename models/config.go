package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (cfg *Config) LoadConfiguration(file string) {
	configFile, err := os.Open(file)
	deferredClose(configFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&cfg)
}

func deferredClose(configFile *os.File) {
	defer configFile.Close()
}

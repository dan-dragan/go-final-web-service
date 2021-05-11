package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (pCfg *Config) LoadConfiguration(file string) error {
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err2 := jsonParser.Decode(pCfg)
	if err2 != nil {
		fmt.Println(err2.Error())
		return err2
	}
	return nil
}

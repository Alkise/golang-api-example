package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config Application configuration
var Config *Configuration

// Configuration Configuration structure
type Configuration struct {
	Database struct {
		Provider string `json:"provider"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// LoadConfiguration Load configuration from json file
func LoadConfiguration(filepath string) (config *Configuration) {
	configFile, err := os.Open(filepath)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

package config

import (
	"encoding/json"
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
func LoadConfiguration(filepath string) (*Configuration, error) {
	configFile, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&Config)

	return Config, nil
}

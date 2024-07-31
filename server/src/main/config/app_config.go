package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	OllamaAPIUrl string `json:"ollamaAPIUrl"`
	AllowedOrigin string `json:"allowedOrigin"`
}

var AppConfig Config

func LoadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return fmt.Errorf("Error decoding config file: %v", err)
	}

	return nil
}


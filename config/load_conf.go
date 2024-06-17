package conf

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB DBConfig `json:"db"`
}

type DBConfig struct {
	DSN string `json:"dsn"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("conf/local.go")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

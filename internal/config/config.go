package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username
	return write(*cfg)
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}

	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	cfg := Config{}
	err = decoder.Decode(&cfg)

	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Your home directory does not exist, or something went wrong: %w", err)
	}
	configPath := filepath.Join(homeDir, configFileName)

	return configPath, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer data.Close()

	encoder := json.NewEncoder(data)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

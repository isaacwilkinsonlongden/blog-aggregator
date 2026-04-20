package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(homeDir, configFileName)
	return configFilePath, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	err := write(*cfg)
	if err != nil {
		return err
	}

	return nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

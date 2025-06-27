package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const cfgFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_username"`
}

func Read() (Config, error) {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(cfgPath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file at %s", cfgPath)
	}
	defer file.Close()

	cfg := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing the configuration located at %s", cfgPath)
	}

	return cfg, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username
	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error fetching home directory")
	}
	fullPath := filepath.Join(homePath, cfgFileName)

	return fullPath, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	rawCfg := Config{
		DBURL:           cfg.DBURL,
		CurrentUsername: cfg.CurrentUsername,
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating/accessing configuration file at: %v", filePath)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(rawCfg)
	if err != nil {
		return fmt.Errorf("error writing the following configuration to file: %v", rawCfg)
	}

	return nil
}

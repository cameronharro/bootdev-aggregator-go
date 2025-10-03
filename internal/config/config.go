package config

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(homeDir, configFileName), nil
}

func Read() (Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	contents, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	result := Config{}
	if err = json.Unmarshal(contents, &result); err != nil {
		return Config{}, err
	}
	return result, nil
}

func write(newConfig []byte) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, newConfig, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	newConfigData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return write(newConfigData)
}

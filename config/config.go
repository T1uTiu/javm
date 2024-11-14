package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	JdkDir string `json:"jdk_dir"`
	EnvVar string `json:"env_var"`
}

func GetConfig() (*Config, error) {
	var config Config

	file, err := os.ReadFile(filepath.Join(os.Getenv("JAVM_HOME"), "config.json"))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

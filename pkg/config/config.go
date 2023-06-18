package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName     string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func ReadConfig(filePath string) (Config, error) {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return config, nil
}

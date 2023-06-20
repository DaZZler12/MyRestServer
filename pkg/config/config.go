package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type JwtConfig struct {
	API_SECRET          string `yaml:"API_SECRET"`
	TOKEN_HOUR_LIFESPAN int    `yaml:"TOKEN_HOUR_LIFESPAN"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	JWT      JwtConfig      `yaml:"jwt"`
}

var (
	Jwtconfig *JwtConfig
	Dbconfig  *DatabaseConfig
)

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
	Jwtconfig = &config.JWT
	Dbconfig = &config.Database
	return config, nil
}

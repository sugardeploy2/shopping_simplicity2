package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}
	Server struct {
		Port int
	}
	// Add other configuration fields as needed
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

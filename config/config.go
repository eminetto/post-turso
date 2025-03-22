package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	DBName           string `mapstructure:"DBNAME"`
	TursoDatabaseURL string `mapstructure:"TURSO_DATABASE_URL"`
	TursoAuthToken   string `mapstructure:"TURSO_AUTH_TOKEN"`
}

func GetConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("parsing config data: %w", err)
	}
	return &config, nil
}

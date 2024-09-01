package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AuthentikURL  string `mapstructure:"AUTHENTIK_URL"`
	ClientID      string `mapstructure:"CLIENT_ID"`
	ClientSecret  string `mapstructure:"CLIENT_SECRET"`
	RedirectURL   string `mapstructure:"REDIRECT_URL"`
	SessionSecret string `mapstructure:"SESSION_SECRET"`
	RedisURL      string `mapstructure:"REDIS_URL"`
}

func Load() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault("AUTHENTIK_URL", "http://localhost:9000")
	viper.SetDefault("REDIRECT_URL", "http://localhost:8080/callback")
	viper.SetDefault("REDIS_URL", "localhost:6379")

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppHost     string `mapstructure:"APP_HOST"`
	AppPort     string `mapstructure:"APP_PORT"`
	PostgresDSN string `mapstructure:"DB_CONNECTION_URL"`
}

func New() Config {
	var config Config

	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}
	viper.SetDefault("APP_HOST", "localhost")
	viper.SetDefault("APP_PORT", "8000")
	viper.Unmarshal(&config)
	return config
}

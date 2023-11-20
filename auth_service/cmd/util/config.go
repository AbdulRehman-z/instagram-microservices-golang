package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER              string        `mapstructure:"DB_DRIVER"`
	DB_URL                 string        `mapstructure:"DB_URL"`
	SYMMETRIC_KEY          string        `mapstructure:"SYMMETRIC_KEY"`
	LISTEN_ADDR            string        `mapstructure:"LISTEN_ADDR"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	DB_MIGRATION_URL       string        `mapstructure:"DB_MIGRATION_URL"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	REDIS_ADDR             string        `mapstructure:"REDIS_ADDR"`
	APP_PASSWORD           string        `mapstructure:"APP_PASSWORD"`
	EMAIL_HOST             string        `mapstructure:"EMAIL_HOST"`
	EMAIL_PORT             string        `mapstructure:"EMAIL_PORT"`
	EMAIL_USERNAME         string        `mapstructure:"EMAIL_USERNAME"`
	EMAIL_FROM             string        `mapstructure:"EMAIL_FROM"`
}

func loadConfig(path string) (*Config, error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

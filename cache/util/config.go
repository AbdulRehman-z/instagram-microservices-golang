package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	LISTEN_ADDR    string `mapstructure:"LISTEN_ADDR"`
	REDIS_ADDR     string `mapstructure:"REDIS_ADDR"`
	APP_PASSWORD   string `mapstructure:"APP_PASSWORD"`
	EMAIL_HOST     string `mapstructure:"EMAIL_HOST"`
	EMAIL_PORT     string `mapstructure:"EMAIL_PORT"`
	EMAIL_USERNAME string `mapstructure:"EMAIL_USERNAME"`
	EMAIL_FROM     string `mapstructure:"EMAIL_FROM"`
	REDIS_HOST     string `mapstructure:"REDIS_HOST"`
	REDIS_PORT     string `mapstructure:"REDIS_PORT"`
	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig(path string) (*Config, error) {

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

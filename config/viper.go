package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	DSN  string `mapstructure:"DSN"`
	Port string `mapstructure:"Port"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			DSN:  os.Getenv("DSN"),
			Port: os.Getenv("Port"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// validate config here

	if config.DSN == "" {
		err = errors.New("DSN is required")
		return
	}

	if config.Port == "" {
		err = errors.New("Port is required")
	}

	return
}

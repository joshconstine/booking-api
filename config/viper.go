package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	DSN  string `mapstructure:"DSN"`
	PORT string `mapstructure:"PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// Fallback to environment variables if config file is not found
		config.DSN = os.Getenv("DSN")
		config.PORT = os.Getenv("PORT")
		if config.DSN == "" || config.PORT == "" {
			return config, fmt.Errorf("error loading config, %v", err)
		}
		return config, nil
	}

	err = viper.Unmarshal(&config)

	// validate config here

	if config.DSN == "" {
		err = errors.New("DSN is required")
		return
	}

	if config.PORT == "" {
		err = errors.New("PORT is required")
	}

	return
}

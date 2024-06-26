package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	DSN                       string `mapstructure:"DSN"`
	PORT                      string `mapstructure:"PORT"`
	OBJECT_STORAGE_URL        string `mapstructure:"OBJECT_STORAGE_URL"`
	OBJECT_STORAGE_ACCESS_KEY string `mapstructure:"OBJECT_STORAGE_ACCESS_KEY"`
	OBJECT_STORAGE_SECRET     string `mapstructure:"OBJECT_STORAGE_SECRET"`
	OBJECT_STORAGE_BUCKET     string `mapstructure:"OBJECT_STORAGE_BUCKET"`
	SEND_GRID_KEY             string `mapstructure:"SEND_GRID_KEY"`
	PAYPAL_CLIENT_ID          string `mapstructure:"PAYPAL_CLIENT_ID"`
	PAYPAL_CLIENT_SECRET      string `mapstructure:"PAYPAL_CLIENT_SECRET"`
}

func LoadConfig(configPath string) (config EnvVars, err error) {
	// viper.AddConfigPath(".")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// Fallback to environment variables if config file is not found
		config.DSN = os.Getenv("DSN")
		config.PORT = os.Getenv("PORT")
		config.OBJECT_STORAGE_URL = os.Getenv("OBJECT_STORAGE_URL")
		config.OBJECT_STORAGE_BUCKET = os.Getenv("OBJECT_STORAGE_BUCKET")
		config.OBJECT_STORAGE_ACCESS_KEY = os.Getenv("OBJECT_STORAGE_ACCESS_KEY")
		config.OBJECT_STORAGE_SECRET = os.Getenv("OBJECT_STORAGE_SECRET")
		config.SEND_GRID_KEY = os.Getenv("SEND_GRID_KEY")
		config.PAYPAL_CLIENT_ID = os.Getenv("PAYPAL_CLIENT_ID")
		config.PAYPAL_CLIENT_SECRET = os.Getenv("PAYPAL_CLIENT_SECRET")

		if config.DSN == "" || config.PORT == "" || config.OBJECT_STORAGE_URL == "" || config.OBJECT_STORAGE_ACCESS_KEY == "" || config.OBJECT_STORAGE_SECRET == "" || config.OBJECT_STORAGE_BUCKET == "" || config.SEND_GRID_KEY == "" || config.PAYPAL_CLIENT_ID == "" || config.PAYPAL_CLIENT_SECRET == "" {
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

	if config.OBJECT_STORAGE_URL == "" {
		err = errors.New("OBJECT_STORAGE_URL is required")

	}

	if config.OBJECT_STORAGE_ACCESS_KEY == "" {
		err = errors.New("OBJECT_STORAGE_ACCESS_KEY is required")
	}

	if config.OBJECT_STORAGE_SECRET == "" {
		err = errors.New("OBJECT_STORAGE_SECRET is required")
	}

	if config.OBJECT_STORAGE_BUCKET == "" {
		err = errors.New("OBJECT_STORAGE_BUCKET is required")
	}

	if config.SEND_GRID_KEY == "" {
		err = errors.New("SEND_GRID_KEY is required")
	}

	if config.PAYPAL_CLIENT_ID == "" {
		err = errors.New("PAYPAL_CLIENT_ID is required")
	}

	if config.PAYPAL_CLIENT_SECRET == "" {
		err = errors.New("PAYPAL_CLIENT_SECRET is required")
	}

	return
}

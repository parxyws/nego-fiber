package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Admin    AdminConfig
}

type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	SSL          bool
	JWTSecretKey string
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	NameDB   string
}

type AdminConfig struct {
	User     string
	Email    string
	Password string
}

func NewAppConfig() (*Config, error) {
	filename := "config.yaml"

	config := viper.New()
	config.SetConfigFile(filename)
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		var FileNotFoundErr viper.ConfigFileNotFoundError
		if errors.As(err, &FileNotFoundErr) {
			return nil, err
		}

		return nil, err
	}

	cfg := new(Config)
	if err := config.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

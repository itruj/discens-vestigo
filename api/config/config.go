package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseName string
}

func loadConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process("dvest", &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()

	if err != nil {
		panic(err)
	}

	return cfg
}

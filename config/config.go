package config

import "os"

type Config struct {
	HOST     string
	PORT     string
	DbConfig *DbConfig
}

func NewConfig() *Config {
	return &Config{
		HOST:     os.Getenv("SERVICE_HOST"),
		PORT:     os.Getenv("SERVICE_PORT"),
		DbConfig: NewDbConfig(),
	}
}

package config

import "os"

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}

package config

import (
	"os"
	"time"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBName     string
	DBPort     string
	DBDriver   string
	DBUser     string
	DBPassword string

	CtxTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_DATABASE"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}

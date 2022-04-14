package config

import (
	"fmt"
	"os"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadConfig() (config Config) {
	config = Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}
	return
}

func (c *Config) GetDBUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DBName)
}

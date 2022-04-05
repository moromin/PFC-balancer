package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load("config/dev.env")
	if err != nil {
		return
	}

	config = Config{
		Port:  os.Getenv("PORT"),
		DBUrl: os.Getenv("DB_URL"),
	}

	return
}

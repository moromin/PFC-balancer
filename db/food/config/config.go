package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

func LoadConfig() Config {
	err := godotenv.Load("config/dev.env")
	if err != nil {
		log.Fatal(err)
	}

	config := Config{
		DBUrl: os.Getenv("DB_URL"),
	}

	return config
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBUrl        string
	JWTSecretKey string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load("config/dev.env")
	if err != nil {
		return
	}

	config = Config{
		Port:         os.Getenv("PORT"),
		DBUrl:        os.Getenv("DB_URL"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}

	return
}

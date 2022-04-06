package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	AuthSvcUrl string
	FoodSvcUrl string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load("config/dev.env")
	if err != nil {
		return
	}

	config = Config{
		Port:       os.Getenv("PORT"),
		AuthSvcUrl: os.Getenv("AUTH_SVC_URL"),
		FoodSvcUrl: os.Getenv("FOOD_SVC_URL"),
	}

	return
}

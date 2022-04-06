package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	FoodSvcUrl string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load("config/dev.env")
	if err != nil {
		return
	}

	config = Config{
		Port:       os.Getenv("PORT"),
		FoodSvcUrl: os.Getenv("FOOD_SVC_URL"),
	}

	return
}

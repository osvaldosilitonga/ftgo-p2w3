package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can't load .env file")
	}
}

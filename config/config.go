package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// load root .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Couldn't load '.env' file")
	}
}

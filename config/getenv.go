package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

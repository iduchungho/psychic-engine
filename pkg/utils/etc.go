package utils

import "github.com/joho/godotenv"

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	return
}

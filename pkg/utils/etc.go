package utils

import (
	"github.com/joho/godotenv"
	"smhome/pkg/repository"
	"strings"
)

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	return
}

func CheckPath(path string) bool {
	for _, p := range repository.DefaultRoutes {
		if path == p {
			return true
		} else if strings.Contains(path, p) {
			return true
		}
	}
	return false
}

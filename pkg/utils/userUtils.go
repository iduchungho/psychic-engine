package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

func GenPassword(pass string) ([]byte, error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		return nil, errEnv
	}
	cost, _ := strconv.Atoi(os.Getenv("COST"))
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return hash, err
}

func ComparePassword(hashPass string, pass string) error {
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	return errCompare
}

func GenerateToken(id string) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 1 day
	})
}

func RequireID(c *fiber.Ctx) (*string, error) {
	id := c.Query("id", "none")
	if id == "none" {
		return nil, errors.New("require ?id = ...")
	}
	return &id, nil
}

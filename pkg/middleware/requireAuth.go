package middleware

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"smhome/config"
	"smhome/app/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireUser(c *fiber.Ctx) error{
	// Get the cookie off req
	tokenString := c.Cookies("Authorization")

	// if err != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// }

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// find the user with token sub
		filter := bson.D{{"id", claims["sub"]}}
		collection := database.GetConnection().Database("SmartHomeDB").Collection("Users")

		var user model.User
		errFind := collection.FindOne(context.TODO(), filter).Decode(&user)

		if errFind != nil {
			// c.AbortWithStatus(http.StatusUnauthorized)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// attach to req
		// cookie := new(fiber.Cookie)
		// cookie.Name = "john"
		// cookie.Value = 
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// c.Cookie("user", user)

		// continue
		c.Next()
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

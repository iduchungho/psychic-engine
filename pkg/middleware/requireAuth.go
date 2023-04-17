package middleware

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"smhome/app/models"
	"smhome/platform/cache"
	"smhome/platform/database"
	"time"
)

func RequireUserID(c *fiber.Ctx) error {
	id := c.Params("id")
	sess, err := cache.GetSessionStoreSlice(id).Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	tkStr := sess.Get("Authorization")
	if tkStr == nil || tkStr == -1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    "no token in session",
			"success": false,
		})
	}

	tokenString := tkStr.(string)

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
		errNext := c.Next()
		if errNext != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	return c.SendStatus(fiber.StatusOK)
}

func RequireUser(c *fiber.Ctx) error {
	// Get the cookie off req
	sess, err := cache.GetSessionStore().Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	tkStr := sess.Get("Authorization")
	if tkStr == nil || tkStr == -1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    "no token in session",
			"success": false,
		})
	}

	tokenString := tkStr.(string)

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
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "not authorized",
				"success": false,
			})
		}
		// attach to req
		// cookie := new(fiber.Cookie)
		// cookie.Name = "john"
		// cookie.Value =
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// c.Cookie("user", user)

		// continue
		errNext := c.Next()
		if errNext != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	return c.SendStatus(fiber.StatusOK)
}

func RequireUserByID(c *fiber.Ctx) error {
	// Get the cookie off req
	sess, err := cache.GetSessionStore().Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	tkStr := sess.Get("Authorization")
	if tkStr == nil || tkStr == -1 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    "no token in session",
			"success": false,
		})
	}

	tokenString := tkStr.(string)

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
		//filter := bson.D{{"id", claims["sub"]}}
		//collection := database.GetConnection().Database("SmartHomeDB").Collection("Users")
		//
		//var user model.User
		//errFind := collection.FindOne(context.TODO(), filter).Decode(&user)

		//if errFind != nil {
		//	// c.AbortWithStatus(http.StatusUnauthorized)
		//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		//		"message": "not authorized",
		//		"success": false,
		//	})
		//}
		if c.Params("id") != claims["sub"] {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "not authorized",
				"success": false,
			})
		}
		// attach to req
		// cookie := new(fiber.Cookie)
		// cookie.Name = "john"
		// cookie.Value =
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// c.Cookie("user", user)

		// continue
		errNext := c.Next()
		if errNext != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	return c.SendStatus(fiber.StatusOK)
}
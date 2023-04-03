package middleware

import (
	"github.com/gofiber/fiber/v2"
	"smhome/pkg/utils"
)

func Redirect(c *fiber.Ctx) error {
	if !utils.CheckPath(c.Path()) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "404 not found",
		})
	}
	return c.Next()
}

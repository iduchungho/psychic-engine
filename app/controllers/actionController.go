package controller

import (
	"github.com/gofiber/fiber/v2"
	service "smhome/pkg/services"
)

func GetActionByID(c *fiber.Ctx) error {
	id := c.Query("id", "none")
	if id == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}
	actionService := service.NewActionService()
	actions, err := actionService.GetActionByUserID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    actions,
		"success": true,
	})
}
func PushActionLog(c *fiber.Ctx) error {
	actionService := service.NewActionService()
	actionOj, err := actionService.Push(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   *actionOj,
		"status": true,
	})

}

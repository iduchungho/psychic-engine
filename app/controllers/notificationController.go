package controller

import (
	"github.com/gofiber/fiber/v2"
	service "smhome/pkg/services"
	"smhome/pkg/utils"
)

func GetNotyByUserID(c *fiber.Ctx) error {
	id, err := utils.RequireID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}
	notyService := service.NewNotifyService()
	res, err := notyService.GetNotifyByUser(*id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
		"success": true,
	})
}

func PushNoty(c *fiber.Ctx) error {
	notyService := service.NewNotifyService()
	res, err := notyService.PushNotify(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
		"success": true,
	})
}

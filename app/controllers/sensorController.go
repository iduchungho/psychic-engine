package controller

import (
	"github.com/gofiber/fiber/v2"
	"smhome/pkg/services"
	"smhome/pkg/utils"
)

func GetTemperature(c *fiber.Ctx) error {
	serviceSensor := service.NewSensorService()
	stat, err := serviceSensor.GetTemperature()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": utils.Statistical(stat),
	})
}

func GetHumidity(c *fiber.Ctx) error {
	serviceSensor := service.NewSensorService()
	stat, err := serviceSensor.GetHumility()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": utils.Statistical(stat),
	})
}

func GetLight(c *fiber.Ctx) error {
	serviceSensor := service.NewSensorService()
	stat, err := serviceSensor.GetLight()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": utils.Statistical(stat),
	})
}

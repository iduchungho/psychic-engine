package controller

import (
	"github.com/gofiber/fiber/v2"
	repo "smhome/pkg/repository"
	service "smhome/pkg/services"
)

func SensorStats(c *fiber.Ctx) error {
	typ := c.Query("type", "none")
	if typ == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?typ = ...",
			"success": false,
		})
	}
	var sensor struct {
		service string
		data    string
	}
	switch typ {
	case "temp":
		sensor.service = repo.DTemp
		sensor.data = repo.TEMPERATURE
	case "humid":
		sensor.service = repo.DHumid
		sensor.data = repo.HUMIDITY
	case "light":
		sensor.service = repo.DLight
		sensor.data = repo.LIGHT
	default:
		return c.SendStatus(fiber.StatusBadRequest)
	}
	dataService := service.NewDataService(sensor.service)
	data, err := dataService.GetSensorData(sensor.data)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": data,
		"success": true,
	})
}

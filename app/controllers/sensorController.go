package controller

import (
	model "smhome/app/models"
	"smhome/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func GetTemperature(c *fiber.Ctx) error {
	nSensors, err := service.NewEntityContext("sensors")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't create sensors models",
		})
	}
	err = nSensors.SetElement("type", "temperature")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, errSen := nSensors.GetEntity("")
	if errSen != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errSen.Error(),
		})
	}

	errIs := nSensors.InsertData(res)
	if errIs != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errIs.Error(),
		})
	}

	stat, ok := res.(model.Sensors)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ok,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": service.Statistical(&stat),
	})
}

func GetHumidity(c *fiber.Ctx) error {
	nSensors, err := service.NewEntityContext("sensors")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't create sensors models",
		})
	}
	err = nSensors.SetElement("type", "humidity")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, errSen := nSensors.GetEntity("")
	if errSen != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errSen.Error(),
		})
	}

	errIs := nSensors.InsertData(res)
	if errIs != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errIs.Error(),
		})
	}

	stat, ok := res.(model.Sensors)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ok,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": service.Statistical(&stat),
	})
}

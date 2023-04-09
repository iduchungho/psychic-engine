package controller

import (
	"github.com/gofiber/fiber/v2"
	model "smhome/app/models"
	"smhome/pkg/repository"
	"smhome/pkg/services"
	"time"
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

	adafruitPkg, errSen := nSensors.GetEntity("")
	if errSen != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errSen.Error(),
		})
	}

	stat, ok := adafruitPkg.(model.Sensors)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ok,
		})
	}

	mongoTempPkg, _ := nSensors.FindDocument("type", "temperature")
	if mongoTempPkg == nil {
		err := nSensors.InsertData(adafruitPkg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t := time.Now().Format(repository.LayoutActionTimestamp)
		err = nSensors.UpdateData("uploaded", t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	} else {
		err := nSensors.UpdateData("payload", stat.Payload)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err = nSensors.UpdateData("edited", stat.Payload[0].CreatedAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t := time.Now().Format(repository.LayoutActionTimestamp)
		err = nSensors.UpdateData("updated", t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": service.Statistical(&stat),
	})
}

func GetHumidity(c *fiber.Ctx) error {
	//humidity
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

	adafruitPkg, errSen := nSensors.GetEntity("")
	if errSen != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errSen.Error(),
		})
	}

	stat, ok := adafruitPkg.(model.Sensors)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ok,
		})
	}

	mongoTempPkg, _ := nSensors.FindDocument("type", "humidity")
	if mongoTempPkg == nil {
		err := nSensors.InsertData(adafruitPkg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t := time.Now().Format(repository.LayoutActionTimestamp)
		err = nSensors.UpdateData("uploaded", t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	} else {
		err := nSensors.UpdateData("payload", stat.Payload)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err = nSensors.UpdateData("edited", stat.Payload[0].CreatedAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t := time.Now().Format(repository.LayoutActionTimestamp)
		err = nSensors.UpdateData("uploaded", t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": service.Statistical(&stat),
	})
}

func GetLight(c *fiber.Ctx) error {
	//humidity
	nSensors, err := service.NewEntityContext("sensors")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't create sensors models",
		})
	}
	err = nSensors.SetElement("type", "light")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	adafruitPkg, errSen := nSensors.GetEntity("")
	if errSen != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errSen.Error(),
		})
	}

	stat, ok := adafruitPkg.(model.Sensors)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": ok,
		})
	}

	mongoTempPkg, _ := nSensors.FindDocument("type", "light")
	if mongoTempPkg == nil {
		err := nSensors.InsertData(adafruitPkg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t := time.Now().Format(repository.LayoutActionTimestamp)
		err = nSensors.UpdateData("uploaded", t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	} else {
		err := nSensors.UpdateData("payload", stat.Payload)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err = nSensors.UpdateData("edited", stat.Payload[0].CreatedAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t := time.Now().Format(repository.LayoutActionTimestamp)
		err = nSensors.UpdateData("uploaded", t)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": service.Statistical(&stat),
	})
}

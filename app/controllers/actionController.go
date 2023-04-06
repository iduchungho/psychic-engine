package controller

import (
	"github.com/gofiber/fiber/v2"
	model "smhome/app/models"
	service "smhome/pkg/services"
)

func GetActionByUsername(c *fiber.Ctx) error {
	return nil
}

func PushActionLog(c *fiber.Ctx) error {
	var body struct {
		UserAction string `json:"user_action"`
		Sensor     string `json:"sensor"`
		ActionName string `json:"action_name"`
		Status     string `json:"status"`
		StatusDesc string `json:"status_desc"`
	}

	if c.BodyParser(&body) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	var actionOj model.Action
	actionOj.UserAction = body.UserAction
	actionOj.Sensor = body.Sensor
	actionOj.ActionName = body.ActionName
	actionOj.StatusDesc = body.StatusDesc
	actionOj.Status = body.Status

	action, _ := service.NewEntityContext("action")
	err := action.InsertData(actionOj)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": action,
		"status":  true,
	})

}

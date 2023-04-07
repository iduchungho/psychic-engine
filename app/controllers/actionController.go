package controller

import (
	"github.com/gofiber/fiber/v2"
	model "smhome/app/models"
	service "smhome/pkg/services"
)

func GetActionByID(c *fiber.Ctx) error {
	action, _ := service.NewEntityContext("action")
	id := c.Params("id")
	actions, err := action.GetEntity(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": actions,
	})
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

	id := c.Params("id")

	var actionOj model.Action
	actionOj.UserAction = body.UserAction
	actionOj.Sensor = body.Sensor
	actionOj.ActionName = body.ActionName
	actionOj.StatusDesc = body.StatusDesc
	actionOj.Status = body.Status
	actionOj.UserID = id

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

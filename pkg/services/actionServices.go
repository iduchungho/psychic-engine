package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	interfaces "smhome/app/interface"
	model "smhome/app/models"
	repo "smhome/pkg/repository"
	"smhome/platform/database"
)

type ActionService struct {
	Factory interfaces.IRepoFactory
}

func NewActionService() *ActionService {
	return &ActionService{
		Factory: NewFactory(database.GetCollection(repo.ACTION)),
	}
}

func (action *ActionService) Push(c *fiber.Ctx) (*model.Action, error) {
	var body struct {
		UserAction string `json:"user_action"`
		Sensor     string `json:"sensor"`
		ActionName string `json:"action_name"`
		Status     string `json:"status"`
		StatusDesc string `json:"status_desc"`
	}

	if c.BodyParser(&body) != nil {
		return nil, errors.New("Failed to parse body")
	}

	id := c.Query("id", "none")
	if id == "none" {
		return nil, errors.New("require ?id = ...")
	}

	var actionOj model.Action
	actionOj.UserAction = body.UserAction
	actionOj.Sensor = body.Sensor
	actionOj.ActionName = body.ActionName
	actionOj.StatusDesc = body.StatusDesc
	actionOj.Status = body.Status
	actionOj.UserID = id

	actionRepo := action.Factory.NewActionRepo()
	res, err := actionRepo.CreateAction(actionOj)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (action *ActionService) GetActionByUserID(userid string) ([]model.Action, error) {
	actionRepo := action.Factory.NewActionRepo()
	actions, err := actionRepo.GetAllAction(userid)
	return actions, err
}

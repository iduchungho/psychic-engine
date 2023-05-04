package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	interfaces "smhome/app/interface"
	model "smhome/app/models"
	repo "smhome/pkg/repository"
	"smhome/pkg/utils"
	"smhome/platform/database"
	"time"
)

type NotifyService struct {
	Factory interfaces.IRepoFactory
}

func NewNotifyService() *NotifyService {
	return &NotifyService{
		Factory: NewFactory(database.GetConnection().Database(repo.DB).Collection(repo.NOTIFY)),
	}
}

func (noty *NotifyService) GetNotifyByUser(id string) ([]model.Notification, error) {
	notyRepo := noty.Factory.NewNotifyRepo()
	return notyRepo.GetAllNotify(id)
}

func (noty *NotifyService) PushNotify(c *fiber.Ctx) (*model.Notification, error) {
	id, err := utils.RequireID(c)
	if err != nil {
		return nil, errors.New("require ?id = ...")
	}
	var body struct {
		Content string `json:"content"`
		Type    string `json:"type"`
	}
	if err = c.BodyParser(&body); err != nil {
		return nil, err
	}
	notyRepo := noty.Factory.NewNotifyRepo()
	var notify = new(model.Notification)
	notify.UserId = *id
	notify.Content = body.Content
	notify.Date = time.Now().Format(repo.LayoutActionTimestamp)
	notify.Type = body.Type
	_, err = notyRepo.CreateNotify(*notify)
	if err != nil {
		return nil, err
	}
	return notify, nil
}

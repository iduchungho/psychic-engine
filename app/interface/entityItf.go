package interfaces

import model "smhome/app/models"

type UserRepo interface {
	GetUserByID(id string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user model.User) (*model.User, error)
	DeleteUserByID(id string) error
	UpdateUser(id string, keyword string, value string) (*model.User, error)
}

type SensorRepo interface {
	GetSensorByName(name string) (*model.Sensors, error)
	DeleteSensor(name string) error
	CreateSensor(sensors interface{}) error
	UpdateSensorByName(name string) error
}

type NotifyRepo interface {
	GetAllNotify() (*[]model.Notification, error)
	CreateNotify(payload model.Notification) error
	DeleteNotifyById(id string) error
}

type ActionRepo interface {
	CreateAction(action model.Action) error
	GetAllAction() ([]model.Action, error)
}

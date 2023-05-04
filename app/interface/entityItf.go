package interfaces

import model "smhome/app/models"

type IUserRepo interface {
	GetUserByID(id string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user model.User) (*model.User, error)
	DeleteUserByID(id string) error
	UpdateUser(id string, keyword string, value string) (*model.User, error)
}

type ISensorRepo interface {
	GetSensorByName(name string) (*model.Sensors, error)
	DeleteSensor(name string) error
	CreateSensor(sensors interface{}) error
	UpdateSensorByName(name string) error
}

type INotifyRepo interface {
	GetAllNotify(userID string) ([]model.Notification, error)
	CreateNotify(payload model.Notification) (*model.Notification, error)
	DeleteNotifyById(id string) error
}

type IActionRepo interface {
	CreateAction(action model.Action) (*model.Action, error)
	GetAllAction(userID string) ([]model.Action, error)
}

type IDataRepo interface {
	PushSensorData(data model.SensorData) (*model.SensorData, error)
	UpdateSensorData(data model.SensorData) (*model.SensorData, error)
}

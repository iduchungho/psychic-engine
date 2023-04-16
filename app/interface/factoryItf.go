package interfaces

import model "smhome/app/models"

// Design Pattern

// Factory method
// Repository pattern
// Singleton pattern
// Decorator pattern

type RepoFactory interface {
	NewUserRepo() UserRepo
	NewActionRepo() ActionRepo
	NewNotifyRepo() NotifyRepo
	NewSensorRepo() SensorRepo
	NewDocumentRepo() DocumentRepo
}

type DocumentRepo interface {
	UserRepo
	ActionRepo
	NotifyRepo
	SensorRepo
	GetSensorAdafruit(name string) (*model.Sensors, error)
}

package interfaces

import model "smhome/app/models"

// Design Pattern

// Factory method
// Repository pattern
// Singleton pattern
// Decorator pattern

type IRepoFactory interface {
	NewUserRepo() IUserRepo
	NewActionRepo() IActionRepo
	NewNotifyRepo() INotifyRepo
	NewSensorRepo() ISensorRepo
	NewDocumentRepo() IDocumentRepo
	NewDataRepo() IDataRepo
}

type IDocumentRepo interface {
	IUserRepo
	IActionRepo
	INotifyRepo
	ISensorRepo
	IDataRepo
	GetSensorAdafruit(name string) (*model.Sensors, error)
}

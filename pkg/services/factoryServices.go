package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	interfaces "smhome/app/interface"
	model "smhome/app/models"
)

type EntityRepoFactory struct {
	collection *mongo.Collection
}

func NewFactory(collection *mongo.Collection) *EntityRepoFactory {
	return &EntityRepoFactory{collection: collection}
}

func (e *EntityRepoFactory) NewUserRepo() interfaces.UserRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewActionRepo() interfaces.ActionRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewNotifyRepo() interfaces.NotifyRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewSensorRepo() interfaces.SensorRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewDocumentRepo() interfaces.DocumentRepo {
	return model.NewDocument(e.collection)
}

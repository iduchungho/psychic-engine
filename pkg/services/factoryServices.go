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
func (e *EntityRepoFactory) NewUserRepo() interfaces.IUserRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewActionRepo() interfaces.IActionRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewNotifyRepo() interfaces.INotifyRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewSensorRepo() interfaces.ISensorRepo {
	return model.NewDocument(e.collection)
}
func (e *EntityRepoFactory) NewDocumentRepo() interfaces.IDocumentRepo {
	return model.NewDocument(e.collection)
}

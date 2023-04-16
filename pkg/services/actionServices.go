package service

import (
	interfaces "smhome/app/interface"
	repo "smhome/pkg/repository"
	"smhome/platform/database"
)

type ActionService struct {
	Factory interfaces.RepoFactory
}

func NewActionService() *ActionService {
	return &ActionService{
		Factory: NewFactory(database.GetCollection(repo.ACTION)),
	}
}

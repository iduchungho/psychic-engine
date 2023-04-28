package service

import (
	interfaces "smhome/app/interface"
	model "smhome/app/models"
	repo "smhome/pkg/repository"
	"smhome/platform/database"
)

type SensorService struct {
	Factory interfaces.IRepoFactory
}

func NewSensorService() *SensorService {
	return &SensorService{
		Factory: NewFactory(database.GetCollection(repo.SENSOR)),
	}
}

func (sen *SensorService) GetTemperature() (*model.Sensors, error) {
	sensor := sen.Factory.NewDocumentRepo()
	adafruit, err := sensor.GetSensorAdafruit(repo.TEMPERATURE)
	if err != nil {
		return nil, err
	}
	return adafruit, nil
}

func (sen *SensorService) GetHumility() (*model.Sensors, error) {
	//sen.Factory = NewFactory(database.GetCollection(repo.SENSOR))
	sensor := sen.Factory.NewDocumentRepo()
	adafruit, err := sensor.GetSensorAdafruit(repo.HUMIDITY)
	if err != nil {
		return nil, err
	}
	return adafruit, nil
}

func (sen *SensorService) GetLight() (*model.Sensors, error) {
	//sen.Factory = NewFactory(database.GetCollection(repo.SENSOR))
	sensor := sen.Factory.NewDocumentRepo()
	adafruit, err := sensor.GetSensorAdafruit(repo.LIGHT)
	if err != nil {
		return nil, err
	}
	return adafruit, nil
}

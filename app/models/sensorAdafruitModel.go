package model

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"os"
	repo "smhome/pkg/repository"
)

// sensor Adafruit
// Design Pattern: Decorator Pattern
// Overwrite function : Get(id string)

// SensorAdafruit
// Param -
// sensor : models sensor
// type : string (temperature, humility, light)
type SensorAdafruit struct {
	Conn   *mongo.Collection
	Extend SensorDocx
}

func (sen SensorAdafruit) GetSensorByName(name string) (*Sensors, error) {
	//TODO implement me
	// get from api
	// save to database
	var api string
	switch name {
	case repo.TEMPERATURE:
		api = os.Getenv("API_TEMP")
	case repo.HUMIDITY:
		api = os.Getenv("API_HUMID")
	case repo.LIGHT:
		api = os.Getenv("API_LIGHT")
	default:
		return nil, errors.New(fmt.Sprintf("no type in entity:%s", name))
	}
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}

	//We Read the response body on the line below.
	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errBody
	}

	errSen := json.Unmarshal(body, &sen.Extend.Data.Payload)
	if errSen != nil {
		return nil, errSen
	}
	err = sen.CreateSensor(sen.Extend.Data)
	if err != nil {
		return nil, err
	}
	sen.Extend.Collection = sen.Conn
	data, _ := sen.Extend.GetSensorByName(name)
	sen.Extend.Data = *data
	return &sen.Extend.Data, nil
}

func (sen SensorAdafruit) DeleteSensor(name string) error {
	return SensorDocx{
		Collection: sen.Conn,
		Data:       sen.Extend.Data,
	}.DeleteSensor(name)
}

func (sen SensorAdafruit) CreateSensor(sensors interface{}) error {
	sens, ok := sensors.(Sensors)
	if !ok {
		return errors.New("require a sensors type")
	}
	return SensorDocx{
		Collection: sen.Conn,
		Data:       sens,
	}.CreateSensor(sens)
}

func (sen SensorAdafruit) UpdateSensorByName(name string) error {
	return SensorDocx{
		Collection: sen.Conn,
		Data:       sen.Extend.Data,
	}.UpdateSensorByName(name)
}

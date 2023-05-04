package model

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
	collection *mongo.Collection
}

func NewDocument(collection *mongo.Collection) *Document {
	return &Document{collection: collection}
}

func (doc Document) GetUserByID(id string) (*User, error) {
	return UserDocx{Collection: doc.collection}.GetUserByID(id)
}

func (doc Document) GetUserByUsername(username string) (*User, error) {
	return UserDocx{Collection: doc.collection}.GetUserByUsername(username)
}

func (doc Document) CreateUser(user User) (*User, error) {
	return UserDocx{Collection: doc.collection, Data: user}.CreateUser(user)
}

func (doc Document) DeleteUserByID(id string) error {
	return UserDocx{Collection: doc.collection}.DeleteUserByID(id)
}

func (doc Document) UpdateUser(id string, keyword string, value string) (*User, error) {
	return UserDocx{Collection: doc.collection}.UpdateUser(id, keyword, value)
}

func (doc Document) CreateAction(action Action) (*Action, error) {
	return ActionDocx{Collection: doc.collection, Data: action}.CreateAction(action)
}

func (doc Document) GetAllAction(userID string) ([]Action, error) {
	return ActionDocx{Collection: doc.collection}.GetAllAction(userID)
}

func (doc Document) GetAllNotify(userID string) ([]Notification, error) {
	return NotifyDocx{Collection: doc.collection}.GetAllNotify(userID)
}

func (doc Document) CreateNotify(payload Notification) (*Notification, error) {
	return NotifyDocx{Collection: doc.collection, Data: payload}.CreateNotify(payload)
}

func (doc Document) DeleteNotifyById(id string) error {
	return NotifyDocx{Collection: doc.collection}.DeleteNotifyById(id)
}

func (doc Document) GetSensorByName(name string) (*Sensors, error) {
	return SensorDocx{Collection: doc.collection}.GetSensorByName(name)
}

func (doc Document) DeleteSensor(name string) error {
	return SensorDocx{Collection: doc.collection}.DeleteSensor(name)
}

func (doc Document) CreateSensor(sensors interface{}) error {
	sensor, ok := sensors.(Sensors)
	if !ok {
		return errors.New("require a Sensors type")
	}
	return SensorDocx{Collection: doc.collection, Data: sensor}.CreateSensor(sensors)
}

func (doc Document) UpdateSensorByName(name string) error {
	return SensorDocx{Collection: doc.collection}.UpdateSensorByName(name)
}

func (doc Document) PushSensorData(data SensorData) (*SensorData, error) {
	return SensorDataDocx{Collection: doc.collection}.PushSensorData(data)
}

func (doc Document) UpdateSensorData(data SensorData) (*SensorData, error) {
	return SensorDataDocx{Collection: doc.collection}.UpdateSensorData(data)
}

func (doc Document) GetSensorAdafruit(name string) (*Sensors, error) {
	return SensorAdafruit{Conn: doc.collection}.GetSensorByName(name)
}

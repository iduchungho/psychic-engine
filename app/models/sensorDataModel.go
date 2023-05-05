package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	repo "smhome/pkg/repository"
	"smhome/platform/database"
	"strconv"
)

type SensorData struct {
	Id      string   `json:"id"`
	Type    string   `json:"type"`
	Date    string   `json:"date"`
	TimeID  string   `json:"timeID"`
	Payload []Sensor `json:"payload"`
}

type SensorDataDocx struct {
	Collection *mongo.Collection
	Data       SensorData
}

func (sensData SensorDataDocx) PushSensorData(data SensorData) (*SensorData, error) {
	var collect string
	switch data.Type {
	case repo.TEMPERATURE:
		collect = repo.DTemp
	case repo.HUMIDITY:
		collect = repo.DHumid
	case repo.LIGHT:
		collect = repo.DLight
	}
	count, _ := database.CountDocuments(database.GetConnection().Database(repo.DB), collect)
	count++
	data.Id = strconv.FormatInt(count, 10)
	_, err := sensData.Collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	sensData.Data = data
	return &sensData.Data, nil
}

func (sensData SensorDataDocx) UpdateSensorData(data SensorData) (*SensorData, error) {
	filter := bson.D{{"id", data.Id}}
	update := bson.M{"$set": bson.M{"date": data.Date, "payload": data.Payload}}
	_, err := sensData.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	sensData.Data = data
	return &sensData.Data, nil
}

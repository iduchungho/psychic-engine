package model

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	repo "smhome/pkg/repository"
	"smhome/platform/database"
	"strconv"
	"time"
)

type Sensor struct {
	Id           string `json:"id"`
	Value        string `json:"value"`
	FeedID       int    `json:"feed_id"`
	FeedKey      string `json:"feed_key"`
	CreatedAt    string `json:"created_at"`
	CreatedEpoch int    `json:"created_epoch"`
	Expiration   string `json:"expiration"`
}

type Sensors struct {
	Id       string   `json:"id"`
	Edited   string   `json:"edited"`
	Created  string   `json:"created"`
	Type     string   `json:"type"`
	Uploaded string   `json:"uploaded"`
	Payload  []Sensor `json:"payload"`
}

type SensorDocx struct {
	Data       Sensors
	Collection *mongo.Collection
}

func (s SensorDocx) GetSensorByName(name string) (*Sensors, error) {
	filter := bson.D{{"type", name}}
	err := s.Collection.FindOne(context.TODO(), filter).Decode(&s.Data)
	if err != nil {
		return nil, err
	}
	return &s.Data, nil
}

func (s SensorDocx) DeleteSensor(name string) error {
	filter := bson.D{{"type", name}}
	_, err := s.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (s SensorDocx) CreateSensor(sensors interface{}) error {
	sens, ok := sensors.(Sensors)
	if !ok {
		return errors.New("require a sensor type")
	}
	typ := sens.Payload[0].FeedKey
	sens.Type = typ
	sens.Uploaded = time.Now().Format(repo.LayoutActionTimestamp)
	sens.Edited = sens.Payload[0].CreatedAt

	count, _ := database.CountDocuments(database.GetConnection().Database(repo.DB), repo.SENSOR)
	count++
	_, err := s.GetSensorByName(typ)
	if err != nil {
		sens.Created = time.Now().Format(repo.LayoutActionTimestamp)
		sens.Id = strconv.Itoa(int(count))
		s.Data = sens
		_, err := s.Collection.InsertOne(context.TODO(), sens)
		if err != nil {
			return err
		}
		return nil
	}
	s.Data = sens
	err = s.UpdateSensorByName(typ)
	if err != nil {
		return err
	}
	return nil
}

func (s SensorDocx) UpdateSensorByName(name string) error {
	filter := bson.D{{"type", name}}
	update := bson.M{"$set": bson.M{"uploaded": s.Data.Uploaded, "edited": s.Data.Edited}}
	_, err := s.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

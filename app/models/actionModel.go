package model

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"smhome/pkg/repository"
	"smhome/platform/database"
	"strconv"
	"time"
)

type Action struct {
	UserAction string `json:"user_action"`
	UserID     string `json:"user_id"`
	Sensor     string `json:"sensor"`
	Id         string `json:"id"`
	ActionName string `json:"action_name"`
	Status     string `json:"status"`
	StatusDesc string `json:"status_desc"`
	TimeStamp  string `json:"time_stamp"`
}

func (a *Action) DeleteEntity(key string, value string) error {
	collection := database.GetCollection("Actions")
	filter := bson.D{{key, value}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (a *Action) GetEntity(param string) (interface{}, error) {
	findOptions := options.Find()
	filter := bson.D{{"userid", param}}
	collection := database.GetCollection("Actions")
	var action []*Action
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem Action
		err = cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		action = append(action, &elem)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	err = cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (a *Action) UpdateData(key string, payload interface{}) error {
	return nil
}

func (a *Action) InsertData(payload interface{}) error {
	action, ok := payload.(Action)
	if !ok {
		return errors.New("InitField: Require a Action")
	}
	collection := database.GetConnection().Database("SmartHomeDB").Collection("Actions")
	count, err := database.CountDocuments(
		database.GetConnection().Database("SmartHomeDB"), "Actions")
	if err != nil {
		return err
	}
	count++
	t := time.Now()

	a.Id = strconv.FormatInt(count, 10)
	a.UserAction = action.UserAction
	a.Status = action.Status
	a.StatusDesc = action.StatusDesc
	a.Sensor = action.Sensor
	a.TimeStamp = t.Format(repository.LayoutActionTimestamp)
	a.ActionName = action.ActionName
	a.UserID = action.UserID

	_, err = collection.InsertOne(context.TODO(), a)
	if err != nil {
		return err
	}
	return nil
}

func (a *Action) SetElement(typ string, value interface{}) error {
	return nil
}

func (a *Action) GetElement(msg string) (*string, error) {
	return nil, nil
}

func (a *Action) FindDocument(key string, val string) (interface{}, error) {

	return nil, nil
}

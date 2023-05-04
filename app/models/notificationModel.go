package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Notification struct {
	Content string `json:"content"`
	Date    string `json:"date"`
	Type    string `json:"type"`
	UserId  string `json:"userid"`
}

type NotifyDocx struct {
	Data       Notification
	Collection *mongo.Collection
}

func (n NotifyDocx) GetAllNotify(userID string) ([]Notification, error) {
	filter := bson.D{{"userid", userID}}
	cur, err := n.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, context.Background())

	var Noty []Notification
	if err = cur.All(context.Background(), &Noty); err != nil {
		return nil, err
	}
	return Noty, nil
}

func (n NotifyDocx) CreateNotify(payload Notification) (*Notification, error) {
	_, err := n.Collection.InsertOne(context.TODO(), payload)
	if err != nil {
		return nil, err
	}
	n.Data = payload
	return &n.Data, nil
}

func (n NotifyDocx) DeleteNotifyById(id string) error {
	//TODO implement me
	panic("implement me")
}

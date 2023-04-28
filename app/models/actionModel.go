package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	repo "smhome/pkg/repository"
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

type ActionDocx struct {
	Data       Action
	Collection *mongo.Collection
}

func (a ActionDocx) CreateAction(action Action) (*Action, error) {
	count, _ := database.CountDocuments(database.GetConnection().Database(repo.DB), repo.ACTION)
	count++
	action.TimeStamp = time.Now().Format(repo.LayoutActionTimestamp)
	action.Id = strconv.FormatInt(count, 10)
	_, err := a.Collection.InsertOne(context.TODO(), action)
	if err != nil {
		return nil, err
	}
	a.Data = action
	return &a.Data, nil
}

func (a ActionDocx) GetAllAction(userID string) ([]Action, error) {
	filter := bson.D{{"userid", userID}}
	cur, err := a.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, context.Background())

	var actions []Action
	if err = cur.All(context.Background(), &actions); err != nil {
		return nil, err
	}
	return actions, nil

}

package model

import "go.mongodb.org/mongo-driver/mongo"

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

func (a ActionDocx) CreateAction(action Action) error {
	//TODO implement me
	panic("implement me")
}

func (a ActionDocx) GetAllAction() ([]Action, error) {
	//TODO implement me
	panic("implement me")
}

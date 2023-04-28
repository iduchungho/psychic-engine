package model

import "go.mongodb.org/mongo-driver/mongo"

type Notification struct {
	Status  string `json:"status"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Type    string `json:"type"`
}

type NotifyDocx struct {
	Data       Notification
	Collection *mongo.Collection
}

func (n NotifyDocx) GetAllNotify() ([]Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotifyDocx) CreateNotify(payload Notification) (*Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotifyDocx) DeleteNotifyById(id string) error {
	//TODO implement me
	panic("implement me")
}

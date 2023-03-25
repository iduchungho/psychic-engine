package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"smhome/platform/database"
)

type User struct {
	Type      string `json:"type" form:"type"`
	Id        string `json:"id" form:"id"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	UserName  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Avatar    string `json:"avatar" form:"avatar"`
}

func (u *User) SetElement(typ string, value interface{}) error {
	switch typ {
	case "type":
		u.Type = value.(string)
		return nil
	}
	return nil
}

func (u *User) GetEntity(param string) (interface{}, error) {
	return nil, nil
}

func (u *User) DeleteEntity(param string) error {
	return nil
}

func (u *User) UpdateData(payload interface{}) error {
	return nil
}

func (u *User) InsertData(payload interface{}) error {
	user, ok := payload.(User)
	if !ok {
		return errors.New("InitField: Require a User")
	}

	u.Type = "user"
	u.Id = user.Id
	u.UserName = user.UserName
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Avatar = user.Avatar
	u.Password = user.Password

	res, _ := u.FindDocument("username", u.UserName)
	if res != nil {
		return errors.New("username already exist")
	}

	collection := database.GetConnection().Database("SmartHomeDB").Collection("Users")

	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil

}

func (u *User) FindDocument(key string, val string) (interface{}, error) {
	filter := bson.D{{key, val}}

	collection := database.GetConnection().Database("SmartHomeDB").Collection("Users")
	var res User
	err := collection.FindOne(context.TODO(), filter).Decode(&res)

	// no documents
	if err != nil {
		return nil, err
	}

	u.Type = res.Type
	u.Id = res.Id
	u.UserName = res.UserName
	u.FirstName = res.FirstName
	u.LastName = res.LastName
	u.Avatar = res.Avatar
	u.Password = res.Password

	return res, nil
}

func (u *User) GetElement(msg string) (*string, error) {
	switch msg {
	case "type":
		return &u.Type, nil
	case "username":
		return &u.UserName, nil
	case "password":
		return &u.Password, nil
	case "id":
		return &u.Id, nil
	case "firstname":
		return &u.FirstName, nil
	case "lastname":
		return &u.LastName, nil
	case "avatar":
		return &u.Avatar, nil
	default:
		return nil, errors.New("no element in user entity")
	}
}

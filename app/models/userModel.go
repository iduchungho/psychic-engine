package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	repo "smhome/pkg/repository"
	"smhome/platform/database"
	"strconv"
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

type UserDocx struct {
	Data       User
	Collection *mongo.Collection
}

func (u UserDocx) GetUserByID(id string) (*User, error) {
	filter := bson.D{{"id", id}}
	err := u.Collection.FindOne(context.TODO(), filter).Decode(&u.Data)
	if err != nil {
		return nil, err
	}
	//return &User{
	//	Type:      repo.USER,
	//	Id:        u.Data.Id,
	//	FirstName: u.Data.FirstName,
	//	LastName:  u.Data.LastName,
	//	UserName:  u.Data.UserName,
	//	Password:  u.Data.Password,
	//	Avatar:    u.Data.Avatar,
	//}, nil
	return &u.Data, nil
}

func (u UserDocx) GetUserByUsername(username string) (*User, error) {
	filter := bson.D{{"username", username}}
	err := u.Collection.FindOne(context.TODO(), filter).Decode(&u.Data)
	if err != nil {
		return nil, err
	}
	//return &User{
	//	Type:      repo.USER,
	//	Id:        u.Data.Id,
	//	FirstName: u.Data.FirstName,
	//	LastName:  u.Data.LastName,
	//	UserName:  u.Data.UserName,
	//	Password:  u.Data.Password,
	//	Avatar:    u.Data.Avatar,
	//}, nil
	return &u.Data, nil
}

func (u UserDocx) CreateUser(user User) (*User, error) {
	count, _ := database.CountDocuments(database.GetConnection().Database(repo.DB), repo.USER)
	count++
	user.Id = strconv.Itoa(int(count))
	user.Type = repo.USER
	_, err := u.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	u.Data = user
	return &u.Data, nil
}

func (u UserDocx) DeleteUserByID(id string) error {
	filter := bson.D{{"id", id}}
	_, err := u.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (u UserDocx) UpdateUser(id string, keyword string, value string) (*User, error) {
	filter := bson.D{{"id", id}}
	update := bson.M{"$set": bson.M{keyword: value}}
	_, err := u.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	user, err := u.GetUserByID(id)
	return user, err
}

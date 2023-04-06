package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func IsCollectionExists(database *mongo.Database, collectionName string) (bool, error) {
	names, err := database.ListCollectionNames(context.Background(), bson.M{"name": collectionName})
	if err != nil {
		return false, err
	}
	for _, name := range names {
		if name == collectionName {
			return true, nil
		}
	}
	return false, nil
}

func CountDocuments(database *mongo.Database, collectionName string) (int64, error) {
	filter := bson.M{} // Lọc tất cả các tài liệu
	count, err := database.Collection(collectionName).CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreateCollection(database *mongo.Database, name string) {
	err := database.CreateCollection(context.Background(), name)
	if err != nil {
		if cmdErr, ok := err.(mongo.CommandError); ok {
			if cmdErr.Code == 48 {
				return
			}
		}
		log.Fatal(err)
	}
}

func IsCollectionEmpty(collection *mongo.Collection) (bool, error) {
	var result bson.M
	err := collection.FindOne(context.Background(), bson.M{}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return true, nil
	} else if err != nil {
		return false, err
	}
	return false, nil
}

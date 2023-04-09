package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"os"
	"smhome/pkg/repository"
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

func (s *Sensors) SetElement(typ string, value interface{}) error {
	switch typ {
	case "type":
		s.Type = value.(string)
		return nil
	case "edited":
		s.Edited = value.(string)
		return nil
	}
	return errors.New("unknown type")
}

func (s *Sensors) GetEntity(param string) (interface{}, error) {
	var api string
	typ, _ := s.GetElement("type")
	switch *typ {
	case "temperature":
		api = os.Getenv("API_TEMP")
	case "humidity":
		api = os.Getenv("API_HUMID")
	case "light":
		api = os.Getenv("API_LIGHT")
	default:
		return nil, errors.New(fmt.Sprintf("no type in entity:%s", *typ))
	}
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}

	//We Read the response body on the line below.
	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errBody
	}

	var sensors Sensors
	errSen := json.Unmarshal(body, &sensors.Payload)
	if errSen != nil {
		return nil, errSen
	}

	s.Created = time.Now().Format(repository.LayoutActionTimestamp)
	s.Payload = sensors.Payload
	s.Type = sensors.Payload[0].FeedKey
	sensors.Type = *typ
	return sensors, nil
}

func (s *Sensors) DeleteEntity(key string, value string) error {
	filter := bson.D{{key, value}}
	collection := database.GetCollection("Sensors")
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sensors) UpdateData(key string, payload interface{}) error {
	collection := database.GetConnection().Database("SmartHomeDB").Collection("Sensors")
	filter := bson.D{{"type", s.Type}}
	update := bson.D{{"$set", bson.D{{key, payload}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sensors) InsertData(payload interface{}) error {
	collection := database.GetConnection().Database("SmartHomeDB").Collection("Sensors")
	typ, _ := s.GetElement("type")
	sensors, ok := payload.(Sensors)
	if !ok {
		return errors.New("InitField: Require a Sensors")
	}
	count, _ := database.CountDocuments(database.GetConnection().Database("SmartHomeDB"), "Sensors")
	count++
	sensors.Type = *typ
	sensors.Id = strconv.Itoa(int(count))
	sensors.Created = s.Created
	sensors.Edited = sensors.Payload[0].CreatedAt
	_, err := collection.InsertOne(context.TODO(), sensors)
	if err != nil {
		return err
	}
	return nil
}
func (s *Sensors) FindDocument(key string, val string) (interface{}, error) {

	collection := database.GetConnection().Database("SmartHomeDB").Collection("Sensors")
	filter := bson.D{{key, val}}
	var res Sensors
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	s.Type = res.Type
	s.Payload = res.Payload
	return res, nil
}

func (s *Sensors) GetElement(msg string) (*string, error) {
	switch msg {
	case "type":
		return &s.Type, nil
	default:
		return nil, errors.New("no element in user entity")
	}
}

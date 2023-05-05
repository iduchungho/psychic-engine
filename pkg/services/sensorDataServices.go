package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"os"
	interfaces "smhome/app/interface"
	model "smhome/app/models"
	repo "smhome/pkg/repository"
	"smhome/pkg/utils"
	"smhome/platform/database"
	"strconv"
	"time"
)

type DSensorService struct {
	Factory interfaces.IRepoFactory
}

func NewDataService(typ string) *DSensorService {
	return &DSensorService{
		Factory: NewFactory(database.GetConnection().Database(repo.DB).Collection(typ)),
	}
}

func (DSens *DSensorService) GetSensorData(typ string) (*model.SensorData, error) {
	dataRepo := DSens.Factory.NewDataRepo()
	var api string
	switch typ {
	case repo.TEMPERATURE:
		api = os.Getenv("API_TEMP")
	case repo.HUMIDITY:
		api = os.Getenv("API_HUMID")
	case repo.LIGHT:
		api = os.Getenv("API_LIGHT")
	default:
		return nil, errors.New(fmt.Sprintf("no type in entity:%s", typ))
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

	sensorData := new(model.SensorData)
	errSen := json.Unmarshal(body, &sensorData.Payload)
	if errSen != nil {
		return nil, errSen
	}
	filter, err := utils.FilterOneDay(sensorData.Payload)
	if err != nil {
		return nil, err
	}
	tm, _ := time.Parse(repo.LayoutTimestamp, filter[0].CreatedAt)
	timeID := fmt.Sprintf("%d%d%d", tm.Year(), tm.Month(), tm.Day())
	sensorData.Type = typ
	sensorData.Date = filter[0].CreatedAt
	sensorData.Payload = filter
	sensorData.TimeID = timeID

	var collect string
	switch typ {
	case repo.TEMPERATURE:
		collect = repo.DTemp
	case repo.HUMIDITY:
		collect = repo.DHumid
	case repo.LIGHT:
		collect = repo.DLight
	}

	count, err := database.CountDocuments(database.GetConnection().Database(repo.DB), collect)
	if err != nil {
		return nil, err
	}

	var dataDB model.SensorData

	id := strconv.FormatInt(count, 10)
	if count == 0 {
		_, err = dataRepo.PushSensorData(*sensorData)
		if err != nil {
			return nil, err
		}
		return sensorData, nil
	}
	//log.Fatal(id)
	err = database.GetCollection(collect).FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&dataDB)
	//return &dataDB, errors.New(collect)
	//log.Fatal(err)
	if err != nil && err == mongo.ErrNoDocuments {
		_, err = dataRepo.PushSensorData(*sensorData)
		if err != nil {
			return nil, err
		}
	} else {
		tDB, _ := time.Parse(repo.LayoutTimestamp, dataDB.Date)
		t, _ := time.Parse(repo.LayoutTimestamp, sensorData.Date)
		if tDB.Day() == t.Day() && tDB.Month() == t.Month() && tDB.Year() == t.Year() {
			sensorData.Payload = append(sensorData.Payload, dataDB.Payload...)
			sensorData = utils.RemoveDuplicates(sensorData.Payload)
			sensorData.Id = id
			sensorData.Type = typ
			sensorData.Date = sensorData.Payload[0].CreatedAt
			sensorData.TimeID = timeID
			_, err = dataRepo.UpdateSensorData(*sensorData)
			if err != nil {
				return nil, err
			}
			return sensorData, nil
		} else {
			_, err = dataRepo.PushSensorData(*sensorData)
			if err != nil {
				return nil, err
			}
		}
	}
	return sensorData, nil
}

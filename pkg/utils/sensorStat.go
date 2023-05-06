package utils

import (
	"github.com/pkg/errors"
	model "smhome/app/models"
	repo "smhome/pkg/repository"
	"strconv"
	"time"
)

func SensorDataStat(sens model.SensorData) (interface{}, error) {
	type payload struct {
		Value string `json:"value"`
		Id    string `json:"id"`
	}
	var data struct {
		Type     string       `json:"type"`
		ValueMax string       `json:"valueMax"`
		CurrTime string       `json:"currTime"`
		TimeID   string       `json:"timeID"`
		Latest   model.Sensor `json:"latest"`
		Payload  []payload    `json:"payload"`
	}
	data.Payload = make([]payload, 24)
	for i := range data.Payload {
		data.Payload[i].Id = "none"
		data.Payload[i].Value = "0"
	}
	data.Type = sens.Type
	data.ValueMax = "0"
	data.TimeID = sens.TimeID
	data.CurrTime = time.Now().Format(repo.LayoutActionTimestamp)
	data.Latest = sens.Payload[0]
	timeHour := time.Now().Hour()
	index := len(sens.Payload) - 1

	for i := 0; i < timeHour; i++ {
		if index < 0 {
			index = 0
		}
		data.Payload[i].Id = sens.Payload[index].Id
		data.Payload[i].Value = sens.Payload[index].Value
		value, _ := strconv.Atoi(sens.Payload[index].Value)
		valueMax, _ := strconv.Atoi(data.ValueMax)
		if valueMax <= value {
			data.ValueMax = sens.Payload[index].Value
		}
		index = index - 2
	}
	return data, nil
}

func FilterOneDay(sen []model.Sensor) ([]model.Sensor, error) {
	var result []model.Sensor
	t := time.Now()
	for _, val := range sen {
		tParse, _ := time.Parse(repo.LayoutTimestamp, val.CreatedAt)
		if tParse.Day() == t.Day() && tParse.Month() == t.Month() && tParse.Year() == t.Year() {
			result = append(result, val)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("no sensor data")
	}
	return result, nil
}

func RemoveDuplicates(slice []model.Sensor) *model.SensorData {
	seen := make(map[string]bool)
	var result = new(model.SensorData)

	for _, val := range slice {
		if _, ok := seen[val.CreatedAt]; !ok {
			seen[val.CreatedAt] = true
			result.Payload = append(result.Payload, val)
		}
	}
	return result
}

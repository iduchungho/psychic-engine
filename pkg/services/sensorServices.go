package service

import (
	"fmt"
	"smhome/app/models"
	"smhome/pkg/repository"
	"sort"
	"strconv"
	"time"
)

func newSensors() *model.Sensors {
	return new(model.Sensors)
}

func Statistical(sen *model.Sensors) interface{} {
	type payload struct {
		Value     string `json:"value"`
		Id        string `json:"id"`
		CreatedAt string `json:"created_at"`
	}
	var data struct {
		Type    string    `json:"type"`
		Average string    `json:"average"`
		UnitOf  string    `json:"unit_of"`
		Payload []payload `json:"payload"`
	}

	var avg float32 = 0.0
	var count float32 = 0.0

	for i, item := range sen.Payload {
		value, err := strconv.ParseFloat(item.Value, 32)
		if err != nil {
			return err
		}
		avg += float32(value)
		count++

		var payload payload
		payload.Id = sen.Payload[i].Id
		payload.CreatedAt = sen.Payload[i].CreatedAt
		payload.Value = sen.Payload[i].Value

		data.Payload = append(data.Payload, payload)
	}

	avg = avg / count

	data.Type = sen.Type
	data.Average = fmt.Sprintf("%f", avg)
	if data.Type == "temperature" {
		data.UnitOf = "Celsius"
	} else if data.Type == "humidity" {
		data.UnitOf = "g/m3"
	} else {
		data.UnitOf = "Lumen"
	}
	return data
}

//func SensorGetTimeCreatedAtMax(sen []model.Sensor) string {
//	_time, _ := time.Parse(repository.LayoutTimestamp, sen[0].CreatedAt)
//	// FIXME: BUG HERE
//	//var temp string
//	var count int = -1
//	for i, item := range sen {
//		local, _ := time.Parse(repository.LayoutTimestamp, item.CreatedAt)
//		if local.After(_time) {
//			_time = local
//			//temp = item.CreatedAt
//			count = i
//			fmt.Println(_time)
//			fmt.Println(local)
//		}
//	}
//	return strconv.FormatInt(int64(count), 10)
//}

func Filter(sen []model.Sensor, condition time.Time) []model.Sensor {
	var result []model.Sensor
	for _, item := range sen {
		local, _ := time.Parse(repository.LayoutTimestamp, item.CreatedAt)
		if local.After(condition) {
			result = append(result, item)
		}
	}
	return result
}

func SortByTime(sen []model.Sensor) []model.Sensor {
	sort.Slice(sen, func(i, j int) bool {
		prev, _ := time.Parse(repository.LayoutTimestamp, sen[i].CreatedAt)
		curr, _ := time.Parse(repository.LayoutTimestamp, sen[j].CreatedAt)
		return prev.After(curr)
	})
	return sen
}

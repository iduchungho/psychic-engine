package service

import (
	"fmt"
	"smhome/app/models"
	"strconv"
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
	} else {
		data.UnitOf = "g/m3"
	}

	return data
}

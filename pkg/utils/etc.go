package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	model "smhome/app/models"
	"smhome/pkg/repository"
	"strconv"
	"strings"
)

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
	return
}

func CheckPath(path string) bool {
	for _, p := range repo.DefaultRoutes {
		if path == p {
			return true
		} else if strings.Contains(path, p) {
			return true
		}
	}
	return false
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
	if data.Type == repo.TEMPERATURE {
		data.UnitOf = "Celsius"
	} else if data.Type == repo.HUMIDITY {
		data.UnitOf = "g/m3"
	} else {
		data.UnitOf = "Lumen"
	}
	return data
}

package utils

import (
	model "smhome/app/models"
	repo "smhome/pkg/repository"
	"time"
)

func SensorDataStat(sens *model.SensorData) (*model.SensorData, error) {
	return nil, nil
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

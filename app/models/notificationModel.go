package model

type Notification struct {
	Status string `json:"status"`
	Content string `json:"content"`
	Date string `json:"date"`
	Type string `json:"type"`
}
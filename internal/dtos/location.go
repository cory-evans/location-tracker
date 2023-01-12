package dtos

import "fmt"

type Location struct {
	Lat       float32           `json:"lat"`
	Lon       float32           `json:"lon"`
	Speed     float32           `json:"speed"`
	Acc       float32           `json:"acc"`
	Timestamp LocationTimestamp `json:"timestamp"`
}

type LocationTimestamp struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

func (lt *LocationTimestamp) GetTimestampAsISO() string {
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", lt.Year, lt.Month, lt.Day, lt.Hour, lt.Minute, lt.Second)
}

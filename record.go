package main

import (
	owm "github.com/briandowns/openweathermap" // "owm" for easier use
	"time"
)

type StationSource int

const (
	// OpenWeatherMap project stations
	OpenWeatherMap StationSource = iota
)

type StationIdenitifer struct {
	Source StationSource
	ID     int
	Name   string
}

type WeatherRecord struct {
	StationID    StationIdenitifer
	TimeRecorded time.Time // Recorded in UTC
	GeoPos       owm.Coordinates
	Main         owm.Main
	Wind         owm.Wind // Speed converted to Knots
	Clouds       owm.Clouds
	Rain         owm.Rain
	Snow         owm.Snow
}

func NewWeatherRecord(stationId StationIdenitifer, current *owm.CurrentWeatherData) *WeatherRecord {
	record := new(WeatherRecord)

	record.StationID = stationId
	record.TimeRecorded = time.Now().UTC()

	record.GeoPos = current.GeoPos
	record.Main = current.Main

	record.Wind.Deg = current.Wind.Deg
	// Convert wind speed from MPH to Knots
	record.Wind.Speed = MphToKnots(current.Wind.Speed)

	record.Clouds = current.Clouds
	record.Rain = current.Rain
	record.Snow = current.Snow

	return record
}

func MphToKnots(mph float64) float64 {
	return (mph * 5280) / 6076
}

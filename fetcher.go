package main

import (
	owm "github.com/briandowns/openweathermap" // "owm" for easier use
	"log"
	"os"
)

const maxNumberOfStations int = 60
const lang string = "en"
const units string = "f"

func getStationList() []string {
	return []string{
		"Seattle",
		"Terre-de-Bas",
	}
}

// getCurrent gets the current weather for the provided
// location in the units provided.
func getCurrentByName(name string) *owm.CurrentWeatherData {
	w, err := owm.NewCurrent(units, lang, os.Getenv("OWM_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}
	w.CurrentByName(name)
	return w
}

func getCurrentForEachStation(stations []string) map[string]*owm.CurrentWeatherData {
	currentsMap := make(map[string]*owm.CurrentWeatherData)

	for _, s := range stations {
		current := getCurrentByName(s)
		currentsMap[s] = current
	}

	return currentsMap
}

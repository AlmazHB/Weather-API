package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"weather-api/internal/db"
	"weather-api/models"
)

type AccuWeatherResponse struct {
	Temperature struct {
		Metric struct {
			Value float64 `json:"Value"`
		} `json:"Metric"`
	} `json:"Temperature"`
}

func StartCollector() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			collect()
			<-ticker.C
		}
	}()
}

func collect() {
	url := "https://dataservice.accuweather.com/currentconditions/v1/349727?apikey=JImYaUOHjUHyxlypqfzOGGIuNmukZy4V"
	r, err := http.Get(url)
	if err != nil {
		log.Println("Failed to fetch data:", err)
		return
	}
	defer r.Body.Close()

	var data []AccuWeatherResponse
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Println("Failed to decode response:", err)
		return
	}

	if len(data) > 0 {
		record := models.WeatherReading{
			Time:        time.Now(),
			Temperature: data[0].Temperature.Metric.Value,
			City:        "Ashgabat",
		}
		db.DB.Create(&record)
		log.Println("Saved City:", record.City, record.Temperature)
	}
}

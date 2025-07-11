package handlers

import (
	"encoding/json"
	"net/http"
	"weather-api/internal/db"
	"weather-api/models"
)

// Hundler get weather

func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	var reading models.WeatherReading
	result := db.DB.Order("time desc").First(&reading)
	if result.Error != nil {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(reading)
}

package main

import (
	"fmt"
	"net/http"
	"weather-api/internal/db"
	"weather-api/internal/handlers"
	"weather-api/internal/services"
)

func main() {

	//Init DB

	db.Init()

	// Start service

	services.StartCollector()

	//Handler

	http.HandleFunc("/weather/current", handlers.GetCurrentWeather)
	fmt.Println("🚀 Server running on :8080")

	http.ListenAndServe(":8080", nil)

}

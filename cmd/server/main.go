package main

import (
	"weather-api/internal/db"
	"weather-api/internal/handlers"
	"weather-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	//Init DB

	db.Init()

	// Start service

	services.StartCollector()

	//Handlers

	r := gin.Default()

	r.GET("/weather/current", handlers.GetCurrent)
	r.GET("/weather/historical", handlers.GetHistorical)
	r.GET("/weather/historical/max", handlers.GetMax)
	r.GET("/weather/historical/min", handlers.GetMin)
	r.GET("/weather/historical/avg", handlers.GetAvg)
	r.GET("/weather/by_time", handlers.GetByTime)
	r.GET("/health", handlers.Health)

}

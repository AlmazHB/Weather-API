package handlers

import (
	"net/http"
	"strconv"
	"time"
	"weather-api/internal/db"
	"weather-api/models"

	"github.com/gin-gonic/gin"
)

// /weather/current — последняя запись
func GetCurrent(c *gin.Context) {
	var reading models.WeatherReading
	result := db.DB.Order("time desc").First(&reading)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no data found"})
		return
	}
	c.JSON(http.StatusOK, reading)
}

// /weather/historical — последние 24 часа
func GetHistorical(c *gin.Context) {
	var readings []models.WeatherReading
	since := time.Now().Add(-24 * time.Hour)
	result := db.DB.Where("time >= ?", since).Order("time").Find(&readings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, readings)
}

// /weather/historical/max — максимум за 24 часа
func GetMax(c *gin.Context) {
	var max float64
	since := time.Now().Add(-24 * time.Hour)
	result := db.DB.Model(&models.WeatherReading{}).
		Where("time >= ?", since).
		Select("MAX(temperature)").
		Scan(&max)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch max temperature"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"max": max})
}

// /weather/historical/min — минимум за 24 часа
func GetMin(c *gin.Context) {
	var min float64
	since := time.Now().Add(-24 * time.Hour)
	result := db.DB.Model(&models.WeatherReading{}).
		Where("time >= ?", since).
		Select("MIN(temperature)").
		Scan(&min)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch min temperature"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"min": min})
}

// /weather/historical/avg — среднее за 24 часа
func GetAvg(c *gin.Context) {
	var avg float64
	since := time.Now().Add(-24 * time.Hour)
	result := db.DB.Model(&models.WeatherReading{}).
		Where("time >= ?", since).
		Select("AVG(temperature)").
		Scan(&avg)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch avg temperature"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"avg": avg})
}

// /weather/by_time?ts=UNIX_TIMESTAMP — ближайшая температура к timestamp
func GetByTime(c *gin.Context) {
	tsStr := c.Query("ts")
	if tsStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing ts parameter"})
		return
	}
	tsInt, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ts parameter"})
		return
	}

	var reading models.WeatherReading
	result := db.DB.Raw(`
		SELECT * FROM weather_readings
		ORDER BY ABS(EXTRACT(EPOCH FROM time) - ?) ASC
		LIMIT 1
	`, tsInt).Scan(&reading)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no data found"})
		return
	}

	c.JSON(http.StatusOK, reading)
}

// /health — статус сервера
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

package models

import (
	"time"
)

type WeatherReading struct {
	ID          uint `gorm:"primaryKey"`
	Time        time.Time
	Temperature float64
	City        string
}

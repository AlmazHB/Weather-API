package db

import (
	"fmt"
	"log"
	"os"
	"weather-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	_ = godotenv.Load()

	dsn := os.Getenv("DB_DSN")
	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.AutoMigrate(&models.WeatherReading{})
	if err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}

	DB = db
	fmt.Println("✅ Connected to DB")
}

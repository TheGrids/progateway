package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectionDataBase() {
	// Checking .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dsn, exist := os.LookupEnv("POSTGRES_CONNECT")
	if !exist {
		log.Printf("Error loading .env file")
	}

	// Database connection
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Adding schema to database
	err = database.AutoMigrate(&User{}, &Tokens{}, &EmailCheck{}, &Test{}, &Answer{}, &Question{}, &UserTestResult{}, &UserAnswer{})
	if err != nil {
		return
	}

	DB = database
}

package config

import (
	"fmt"
	"gin-post-api/src/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf("host=localhost user=postgres password=4321 dbname=goTest port=5432 sslmode=disable TimeZone=Asia/Kolkata")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database connected successfully")

	DB.AutoMigrate(&models.User{})
}

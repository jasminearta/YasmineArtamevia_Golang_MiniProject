package config

import (
	"YasmineArtamevia_Golang_MiniProject/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := "root:@tcp(localhost:3306)/miniproject"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.User{}, &models.ProductLog{})
	DB = db
	log.Println("Database connected successfully.")
}

package config

import (
	"fmt"
	"jessie_miniproject/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConfigDB holds database configuration values.
type ConfigDB struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

var DB *gorm.DB

func InitDB() error {
	configDB := ConfigDB{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.User,
		configDB.Password,
		configDB.Host,
		configDB.Port,
		configDB.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&models.ProductLog{}, &models.User{}); err != nil {
		return fmt.Errorf("failed to migrate database models: %w", err)
	}

	DB = db
	return nil
}

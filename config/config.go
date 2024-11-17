package config

import (
	"fmt"
	"jessie_miniproject/models"
	"log"
	"os"

	"github.com/joho/godotenv"
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

// DB is a global GORM database connection.
var DB *gorm.DB

// LoadEnv loads environment variables from the .env file.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded, using system environment variables.")
	}
}

// InitDB initializes the database connection.
func InitDB() error {
	LoadEnv()
	configDB := ConfigDB{
		Host:     getEnv("DATABASE_HOST"),
		User:     getEnv("DATABASE_USER"),
		Password: getEnvOrDefault("DATABASE_PASSWORD", ""), // Default password is empty
		Port:     getEnv("DATABASE_PORT"),
		Name:     getEnv("DATABASE_NAME"),
	}

	// Create DSN (Data Source Name) for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.User,
		configDB.Password,
		configDB.Host,
		configDB.Port,
		configDB.Name)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	log.Println("Database connection established successfully.")

	// Auto-migrate models
	if err := db.AutoMigrate(&models.ProductLog{}, &models.User{}); err != nil {
		return fmt.Errorf("failed to migrate database models: %w", err)
	}
	log.Println("Database migration completed successfully.")

	DB = db
	return nil
}

// getEnv retrieves required environment variables or exits if not set.
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

// getEnvOrDefault retrieves environment variables or uses the default value if not set.
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

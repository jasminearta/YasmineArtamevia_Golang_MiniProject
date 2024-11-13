package main

import (
	"jessie_miniproject/config"
	"jessie_miniproject/controllers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	loadEnv()

	config.InitDB()

	c := echo.New()
	c.Use(middleware.Logger())

	// routes auth
	c.POST("/api/register", controllers.Registrasi)
	c.POST("/api/login", controllers.Login)
	c.GET("/api/logout", controllers.Logout)

	// Start server
	c.Start(":8080")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed load env")
	}
	// Cek apakah JWT_SECRET sudah termuat
	log.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))
}

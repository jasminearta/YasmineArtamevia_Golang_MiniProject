package main

import (
	"YasmineArtamevia_Golang_MiniProject/config"
	"YasmineArtamevia_Golang_MiniProject/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize database
	config.InitDB()

	// Initialize routes
	routes.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

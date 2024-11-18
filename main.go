package main

import (
	"jessie_miniproject/config"
	"jessie_miniproject/controllers"
	middlewares "jessie_miniproject/middlewares"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Memuat file .env

	// Inisialisasi database
	if err := config.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Inisialisasi Echo
	c := echo.New()
	c.Use(middleware.Logger())

	// routes auth
	c.POST("/api/register", controllers.Registrasi)
	c.POST("/api/login", controllers.Login)
	c.GET("/api/logout", controllers.Logout)

	// Protected routes group
	eAuth := c.Group("/api/products")

	// Middleware untuk autentikasi JWT
	eAuth.Use(middlewares.JWTMiddleware)

	// routes produk
	eAuth.POST("", controllers.AddProduct)       // Add new product
	eAuth.GET("", controllers.GetAllProducts)    // Get all products
	eAuth.GET("/:id", controllers.GetByID)       // Get product by ID
	eAuth.PUT("/:id", controllers.UpdateProduct) // Update product by ID
	eAuth.DELETE("/:id", controllers.DeleteProduct)

	// Mulai server
	c.Logger.Fatal(c.Start(":8000"))

}

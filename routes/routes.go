package routes

import (
	"YasmineArtamevia_Golang_MiniProject/controllers"
	"YasmineArtamevia_Golang_MiniProject/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// Authentication routes
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Product management routes (protected by JWT middleware)
	product := api.Group("/products", middlewares.JWTMiddleware)
	product.POST("", controllers.AddProduct)
	product.GET("", controllers.GetAllProducts)

	// Analysis routes
	api.POST("/analyze", controllers.AnalyzeConsumption)
}

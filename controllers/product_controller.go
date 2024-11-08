package controllers

import (
	"YasmineArtamevia_Golang_MiniProject/config"
	"YasmineArtamevia_Golang_MiniProject/models"
	"YasmineArtamevia_Golang_MiniProject/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddProduct(c echo.Context) error {
	var product models.ProductLog
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not add product"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product added successfully"))
}

func GetAllProducts(c echo.Context) error {
	var products []models.ProductLog
	result := config.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not fetch products"))
	}
	return c.JSON(http.StatusOK, products)
}

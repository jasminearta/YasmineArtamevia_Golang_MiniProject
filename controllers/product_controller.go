package controllers

import (
	"jessie_miniproject/config"
	"jessie_miniproject/models"
	"jessie_miniproject/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddProduct(c echo.Context) error {
	var product models.ProductLog

	userID, _ := c.Get("user_id").(int)

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	product.UserID = userID
	result := config.DB.Create(&product)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not add product"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product added successfully"))
}

func GetAllProducts(c echo.Context) error {
	var products []models.ProductLog

	c.Bind(&products)
	result := config.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not fetch products"))
	}
	return c.JSON(http.StatusOK, products)

}

func GetByID(c echo.Context) error {
	var products []models.ProductLog
	id := c.Param("id")
	result := config.DB.First(&products, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not fetch products"))
	}
	return c.JSON(http.StatusOK, products)

}

// controllers/product_controller.go
func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.ProductLog
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	result := config.DB.Model(&models.ProductLog{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not update product"))
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Product not found"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product updated successfully"))
}

// controllers/product_controller.go
func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.ProductLog
	result := config.DB.Delete(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not delete product"))
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Product not found"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product deleted successfully"))
}

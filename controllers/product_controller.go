package controllers

import (
	"context"
	"jessie_miniproject/config"
	"jessie_miniproject/helper"
	"jessie_miniproject/models"
	"jessie_miniproject/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddProduct(c echo.Context) error {
	var product models.ProductLog

	// Get user ID from context and ensure the type is correct
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("User ID tidak valid"))
	}

	// Bind the request payload to the product model
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	product.UserID = userID

	// Generate the AI query and get a description
	query := helper.GenerateProductQuery(product.ProductName, product.Material, product.IsPlastic)
	ctx := context.Background()
	aiResponse, err := helper.ResponseAI(ctx, query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to generate product description"))
	}

	// Set the generated AI description to the product
	product.Rekomendasi = aiResponse

	// Save the product to the database with the AI description
	if result := config.DB.Create(&product); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not add product"))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product added successfully with AI description", product))
}

func GetAllProducts(c echo.Context) error {
	var products []models.ProductLog

	// Mengambil semua produk dari database
	if result := config.DB.Find(&products); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not fetch products"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Products fetched successfully", products))
}

func GetByID(c echo.Context) error {

	var product models.ProductLog

	// Konversi `id` dari string ke uint
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid product ID format"))
	}

	// Mengambil produk berdasarkan ID
	result := config.DB.First(&product, uint(id))
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Product not found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not fetch product: "+result.Error.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product fetched successfully", product))
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.ProductLog

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	// Memperbarui produk di database
	result := config.DB.Model(&models.ProductLog{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not update product"))
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Product not found"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product updated successfully", product))
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.ProductLog

	// Menghapus produk berdasarkan ID
	result := config.DB.Delete(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not delete product"))
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse("Product not found"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product deleted successfully", nil))
}

// AddProductWithAI handles adding a product and generating an AI description
func AddProductWithAI(c echo.Context) error {
	var product models.ProductLog
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	// Generate query for AI response based on product details
	query := helper.GenerateProductQuery(product.ProductName, product.Material, product.IsPlastic)

	// Get AI response
	ctx := context.Background()
	aiResponse, err := helper.ResponseAI(ctx, query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to generate product description"))
	}

	// Set AI description in product log
	product.Rekomendasi = aiResponse // assuming you have a Description field in ProductLog model

	// Save product to the database
	if result := config.DB.Create(&product); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not add product"))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Product added successfully with AI description", product))
}

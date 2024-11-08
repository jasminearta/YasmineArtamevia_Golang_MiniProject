package controllers

import (
	"YasmineArtamevia_Golang_MiniProject/config"
	"YasmineArtamevia_Golang_MiniProject/models"
	"YasmineArtamevia_Golang_MiniProject/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func AnalyzeConsumption(c echo.Context) error {
	var analysis models.PlasticAnalysis
	if err := c.Bind(&analysis); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	analysis.TimeAnalysis = time.Now()

	result := config.DB.Create(&analysis)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not perform analysis"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Analysis added successfully"))
}

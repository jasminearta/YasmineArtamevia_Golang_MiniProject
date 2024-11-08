package controllers

import (
	"YasmineArtamevia_Golang_MiniProject/config"
	"YasmineArtamevia_Golang_MiniProject/models"
	"YasmineArtamevia_Golang_MiniProject/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not register user"))
	}
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("User registered successfully"))
}

func Login(c echo.Context) error {
	var input models.User
	var user models.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload"))
	}

	config.DB.Where("username = ?", input.Username).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, utils.NewErrorResponse("Invalid username or password"))
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Could not generate token"))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

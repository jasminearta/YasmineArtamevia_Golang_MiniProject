package controllers

import (
	"jessie_miniproject/config"
	"jessie_miniproject/model"
	"jessie_miniproject/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "Gagal Mengambil Data",
		})
	}
	oldPassword := user.Password
	if err := config.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return c.JSON(401, map[string]interface{}{
			"message": "Email Tidak Ditemukan",
		})
	}
	if !CheckPasswordHash(oldPassword, user.Password) {
		return c.JSON(401, map[string]interface{}{
			"message": "Password Salah",
		})
	}
	token, _ := GenerateJWT(user.Email, user.ID)

	// set cokie
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Login",
		"token":   token,
	})
}
func Registrasi(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "Gagal Mengambil Data",
		})
	}
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": "Gagal Mendaftarkan User",
		})
	}
	return c.JSON(201, map[string]interface{}{
		"message": "Berhasil Mendaftarkan User",
	})
}
func Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Logout",
	})
}

// hasing dan cek password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generate token
func GenerateJWT(email string, id int) (string, error) {
	claims := &model.JwtCustomClaims{
		UserID: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}

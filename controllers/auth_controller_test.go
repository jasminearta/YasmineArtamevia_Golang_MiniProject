// controllers/auth_test.go
package controllers_test

import (
	"fmt"
	"jessie_miniproject/controllers"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Helper function untuk generate email dan username unik
func generateRandomEmailAndUsername() (string, string) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100000)
	email := fmt.Sprintf("user%d@example.com", randomNumber)
	username := fmt.Sprintf("user%d", randomNumber)
	return email, username
}

func TestRegistrasi_Success(t *testing.T) {
	e := echo.New()

	// Generate email dan username secara dinamis
	email, username := generateRandomEmailAndUsername()

	// Payload registrasi dengan email dan username yang di-generate
	registrationPayload := fmt.Sprintf(`{
		"email": "%s", 
		"username": "%s",
		"password": "newpassword123"
	}`, email, username)

	req := httptest.NewRequest(http.MethodPost, "/registrasi", strings.NewReader(registrationPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Memanggil controller Registrasi
	err := controllers.Registrasi(c)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validasi respons
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "Berhasil Mendaftarkan User")
}

func TestLogout_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Memanggil controller Logout
	err := controllers.Logout(c)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validasi respons
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Berhasil Logout")

	// Validasi cookie
	cookie := rec.Result().Cookies()
	assert.Equal(t, "token", cookie[0].Name)
	assert.Equal(t, "", cookie[0].Value)
	assert.Equal(t, -1, cookie[0].MaxAge)
}

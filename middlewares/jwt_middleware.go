package midlewares

import (
	"jessie_miniproject/models"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// jwt midleware
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h, err := c.Cookie("token")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token not found")
		}

		tokenString := h.Value
		claims := &models.JwtCustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// Tambahkan log ini untuk debugging
		if err != nil {
			c.Logger().Error("Error parsing token:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		if !token.Valid {
			c.Logger().Error("Token is not valid")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user_id", claims.UserID)
		return next(c)
	}
}

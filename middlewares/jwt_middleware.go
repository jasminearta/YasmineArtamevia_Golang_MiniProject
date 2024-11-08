package middlewares

import (
	"YasmineArtamevia_Golang_MiniProject/utils"

	"github.com/labstack/echo/v4/middleware"
)

var JWTMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(utils.GetJWTSecret()),
})

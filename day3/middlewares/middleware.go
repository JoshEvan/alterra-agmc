package middlewares

import (
	"github.com/JoshEvan/alterra-agmc-day3/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitAuth(e *echo.Echo) (eAuth *echo.Group) {
	eAuth = e.Group("")
	config := middleware.JWTConfig{
		ContextKey: "user", TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme: "Bearer",
		SigningKey: []byte(config.SECRET_JWT),
	}
	eAuth.Use(middleware.JWTWithConfig(config))
	return
}

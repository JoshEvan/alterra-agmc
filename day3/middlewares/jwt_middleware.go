package middlewares

import (
	"time"

	"github.com/JoshEvan/alterra-agmc-day3/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const ttl = time.Hour

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractTokenUserID(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		return claims["user_id"].(uint)
	}
	return 0
}

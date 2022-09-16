package middlewares

import (
	"github.com/JoshEvan/alterra-agmc-day3/config"
	"github.com/JoshEvan/alterra-agmc-day3/models"
	"github.com/labstack/echo/v4"
)

func BasicAuthFromDB(username, password string, c echo.Context) (bool, error) {
	if err := config.DB.Where("email = ? AND password = ?", username, password).First(&models.User{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

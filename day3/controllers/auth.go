package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JoshEvan/alterra-agmc-day3/lib/database"
	"github.com/JoshEvan/alterra-agmc-day3/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	userDB, e := database.LoginUser(user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	resp := response{Status: http.StatusText(http.StatusOK), Data: userDB}

	return c.JSON(http.StatusOK, resp)
}

func auth(c echo.Context, userID uint) error {
	authInfo := c.Get("user")
	if authInfo == nil {
		return errors.New("no auth")
	}
	jwtClaims := authInfo.(*jwt.Token).Claims.(jwt.MapClaims)
	authUserID := fmt.Sprintf("%v", jwtClaims["user_id"])
	id := c.Param("id")
	if authUserID != id {
		return errors.New("not eligible")
	}
	return nil
}

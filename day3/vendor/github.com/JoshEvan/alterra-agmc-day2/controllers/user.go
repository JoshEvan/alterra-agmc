package controllers

import (
	"net/http"
	"strconv"

	"github.com/JoshEvan/alterra-agmc-day3/lib/database"
	"github.com/JoshEvan/alterra-agmc-day3/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	books, err := database.GetUsers()
	if err != nil {
		resp := response{Status: http.StatusText(http.StatusInternalServerError), Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := response{Status: http.StatusText(http.StatusOK), Data: books}

	return c.JSON(http.StatusOK, resp)
}

func GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Status: http.StatusText(http.StatusBadRequest), Error: "invalid id"})
	}

	user, err := database.GetUserByID(id)
	resp := response{Status: http.StatusText(http.StatusOK)}
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Data = user
	return c.JSON(http.StatusOK, resp)
}

func CreateUser(c echo.Context) error {
	newUser := models.User{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Password: c.FormValue("password"),
	}

	err := database.AddUserByID(newUser)
	if err != nil {
		resp := response{Status: http.StatusText(http.StatusInternalServerError)}
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := response{Status: http.StatusText(http.StatusOK)}

	return c.JSON(http.StatusOK, resp)
}

func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Status: http.StatusText(http.StatusBadRequest), Error: "invalid id"})
	}
	updatedUserData := models.User{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Password: c.FormValue("password"),
	}

	err = database.UpdateUserByID(id, updatedUserData)
	if err != nil {
		resp := response{Status: http.StatusText(http.StatusInternalServerError)}
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := response{Status: http.StatusText(http.StatusOK)}

	return c.JSON(http.StatusOK, resp)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Status: http.StatusText(http.StatusBadRequest), Error: "invalid id"})
	}

	if err := database.DeleteUserByID(id); err != nil {
		resp := response{Status: http.StatusText(http.StatusInternalServerError)}
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp := response{Status: http.StatusText(http.StatusOK)}

	return c.JSON(http.StatusOK, resp)
}

package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/JoshEvan/alterra-agmc-day3/lib/database"
	"github.com/JoshEvan/alterra-agmc-day3/models"
	"github.com/labstack/echo/v4"
)

type response struct {
	Status string      `json:"status"`
	Error  string      `json:"error_message"`
	Data   interface{} `json:"data"`
}

func GetBooks(c echo.Context) error {
	books := database.GetStaticBooks()
	resp := response{Status: http.StatusText(http.StatusOK), Data: books}

	return c.JSON(http.StatusOK, resp)
}

func GetBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Status: http.StatusText(http.StatusBadRequest), Error: "invalid id"})
	}

	book := database.GetStaticBookByID(id)
	resp := response{Status: http.StatusText(http.StatusOK)}
	if book == nil {
		resp.Error = "data not found"
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp.Data = book
	return c.JSON(http.StatusOK, resp)
}

func CreateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	newBook := models.Book{
		ID:        id,
		Title:     c.FormValue("title"),
		ISBN:      c.FormValue("isbn"),
		Writer:    c.FormValue("writer"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := database.AddStaticBookByID(newBook)
	if err != nil {
		resp := response{Status: http.StatusText(http.StatusBadRequest)}
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := response{Status: http.StatusText(http.StatusOK)}

	return c.JSON(http.StatusOK, resp)
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Status: http.StatusText(http.StatusBadRequest), Error: "invalid id"})
	}

	book := database.GetStaticBookByID(id)
	updatedBookData := models.Book{
		ID:        book.ID,
		Title:     c.FormValue("title"),
		ISBN:      c.FormValue("isbn"),
		Writer:    c.FormValue("writer"),
		CreatedAt: book.CreatedAt,
		UpdatedAt: time.Now(),
	}

	book, err = database.UpdateStaticBookByID(id, updatedBookData)
	if err != nil {
		resp := response{Status: http.StatusText(http.StatusBadRequest)}
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := response{Status: http.StatusText(http.StatusOK), Data: book}

	return c.JSON(http.StatusOK, resp)
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Status: http.StatusText(http.StatusBadRequest), Error: "invalid id"})
	}

	if err := database.DeleteStaticBookByID(id); err != nil {
		resp := response{Status: http.StatusText(http.StatusBadRequest)}
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := response{Status: http.StatusText(http.StatusOK)}

	return c.JSON(http.StatusOK, resp)
}

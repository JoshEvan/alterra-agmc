package database

import (
	"errors"
	"time"

	"github.com/JoshEvan/alterra-agmc-day4/models"
)

func getStaticTime() time.Time {
	staticTime, _ := time.Parse("2006-01-02T15:04:05", "2019-02-18T00:00:00")
	return staticTime
}

var staticBooks = map[int]models.Book{
	1234: {
		ID:        1234,
		Title:     "Judul Buku Satu",
		ISBN:      "1-234-5678-9101112-13",
		Writer:    "Dr. Who",
		CreatedAt: getStaticTime(),
		UpdatedAt: getStaticTime(),
		DeletedAt: getStaticTime(),
	},
}

func DeleteStaticBookByID(id int) error {
	for _, book := range staticBooks {
		if book.ID == id {
			delete(staticBooks, id)
			return nil
		}
	}

	return errors.New("data not found")
}

func UpdateStaticBookByID(id int, updatedBook models.Book) (*models.Book, error) {
	if _, ok := staticBooks[id]; !ok {
		return nil, errors.New("data not found")
	}
	staticBooks[id] = updatedBook
	return &updatedBook, nil
}

func AddStaticBookByID(newBook models.Book) (err error) {
	if _, ok := staticBooks[newBook.ID]; ok {
		return errors.New("duplicate id")
	}
	staticBooks[newBook.ID] = newBook
	return nil
}

func GetStaticBookByID(id int) *models.Book {
	for _, book := range staticBooks {
		if book.ID == id {
			return &book
		}
	}

	return nil
}

func GetStaticBooks() (books []models.Book) {
	for _, v := range staticBooks {
		books = append(books, v)
	}
	return
}

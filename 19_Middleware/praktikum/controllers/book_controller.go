package controllers

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/rest-orm/lib/database"
	"github.com/arvinpaundra/rest-orm/models"
	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))

	book, err := database.GetBook(uint(bookId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	var book models.Book

	c.Bind(&book)

	bookCreated, err := database.CreateBook(book)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create book",
		"book":    bookCreated,
	})
}

func UpdateBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))

	var book models.Book

	c.Bind(&book)

	bookCreated, err := database.UpdateBook(book, uint(bookId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    bookCreated,
	})
}

func DeleteBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))

	_, err := database.DeleteBook(uint(bookId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success delete book",
	})
}

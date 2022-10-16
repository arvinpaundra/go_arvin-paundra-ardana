package controllers

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/unit-test/config"
	"github.com/arvinpaundra/unit-test/models"
	"github.com/labstack/echo/v4"
)

type DTOBook struct {
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}

func GetBooksController(c echo.Context) error {
	books := []DTOBook{}

	err := config.DB.Model(&models.Book{}).Find(&books).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}

func GetBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))

	book := DTOBook{}

	err := config.DB.Model(&models.Book{}).Where("id = ?", bookId).Take(&book).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func CreateBookController(c echo.Context) error {
	book := DTOBook{}

	c.Bind(&book)

	err := config.DB.Model(&models.Book{}).Save(&models.Book{Title: book.Title, Author: book.Author, Publisher: book.Publisher}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"book": book,
	})
}

func UpdateBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))

	book := DTOBook{}

	c.Bind(&book)

	err := config.DB.Model(&models.Book{}).Where("id = ?", bookId).Updates(models.Book{Title: book.Title, Author: book.Author, Publisher: book.Publisher}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func DeleteBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Unscoped().Model(&models.Book{}).Where("id = ?", bookId).Delete(&models.Book{}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success delete book",
	})
}

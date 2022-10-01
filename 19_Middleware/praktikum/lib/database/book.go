package database

import (
	"github.com/arvinpaundra/rest-orm/config"
	"github.com/arvinpaundra/rest-orm/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Model(&models.Book{}).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func GetBook(id uint) (interface{}, error) {
	var book models.Book

	if err := config.DB.Model(&models.Book{}).Where("id = ?", id).Take(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func CreateBook(book models.Book) (interface{}, error) {
	if err := config.DB.Model(&models.Book{}).Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBook(book models.Book, id uint) (interface{}, error) {
	err := config.DB.Model(&models.Book{}).Where("id = ?", id).Updates(models.Book{Title: book.Title, Author: book.Author, Publisher: book.Publisher}).Error

	if err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(id uint) (interface{}, error) {
	err := config.DB.Unscoped().Model(&models.Book{}).Where("id = ?", id).Delete(&models.Book{}).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}

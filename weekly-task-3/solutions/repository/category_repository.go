package repository

import (
	"github.com/arvinpaundra/golang-blog/database"
	"github.com/arvinpaundra/golang-blog/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category model.Category) error
	FindAll() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) Create(category model.Category) error {
	err := database.DB.Model(&model.Category{}).Create(&category).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category

	err := database.DB.Model(&model.Category{}).Find(&categories).Error

	if err != nil {
		return []model.Category{}, err
	}

	return categories, nil
}

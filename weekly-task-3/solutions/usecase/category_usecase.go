package usecase

import (
	"github.com/arvinpaundra/golang-blog/dto"
	"github.com/arvinpaundra/golang-blog/model"
	"github.com/arvinpaundra/golang-blog/repository"
	"github.com/google/uuid"
	"time"
)

type CategoryUsecase interface {
	Create(categoryDTO dto.CategoryDTO) (model.Category, error)
	FindAll() ([]model.Category, error)
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepo}
}

func (c *categoryUsecase) Create(categoryDTO dto.CategoryDTO) (model.Category, error) {
	id := uuid.NewString()
	createdAt := time.Now()
	updatedAt := time.Now()

	category := model.Category{
		ID:           id,
		CategoryName: categoryDTO.CategoryName,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	err := c.categoryRepository.Create(category)

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (c *categoryUsecase) FindAll() ([]model.Category, error) {
	categories, err := c.categoryRepository.FindAll()

	if err != nil {
		return []model.Category{}, err
	}

	return categories, nil
}

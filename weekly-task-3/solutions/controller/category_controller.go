package controller

import (
	"github.com/arvinpaundra/golang-blog/dto"
	"github.com/arvinpaundra/golang-blog/model"
	"github.com/arvinpaundra/golang-blog/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type categoryController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryController(categoryUsecase usecase.CategoryUsecase) *categoryController {
	return &categoryController{categoryUsecase}
}

func (cat *categoryController) HandlerAddCategory(c echo.Context) error {
	categoryDTO := dto.CategoryDTO{}

	if err := c.Bind(&categoryDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Isi data dengan lengkap.",
			"data":    nil,
		})
	}

	category, err := cat.categoryUsecase.Create(categoryDTO)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "Berhasil menambahkan kategori.",
		"data": map[string]model.Category{
			"category": category,
		},
	})
}

func (cat *categoryController) HandlerFindAll(c echo.Context) error {
	categories, err := cat.categoryUsecase.FindAll()

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "sukses get data semua kategori",
		"data": map[string][]model.Category{
			"categories": categories,
		},
	})
}

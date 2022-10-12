package controllers

import (
	"net/http"
	"time"

	"github.com/arvinpaundra/weekly-task-2/databases"
	"github.com/arvinpaundra/weekly-task-2/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetKategorisController(c echo.Context) error {
	kategoris, err := databases.GetKategoris()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get all categories",
		"data": map[string]interface{}{
			"categories": kategoris,
		},
	})
}

func GetKategoriController(c echo.Context) error {
	kategoriId := c.Param("kategoriId")

	kategori, err := databases.GetKategori(kategoriId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get category",
		"data": map[string]interface{}{
			"category": kategori,
		},
	})
}

func CreatekategoriController(c echo.Context) error {
	ID := uuid.NewString()
	CreatedAt := time.Now()
	UpdatedAt := time.Now()

	var kategori models.Kategori

	kategori = models.Kategori{
		ID:        ID,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}

	c.Bind(&kategori)

	newKategori, err := databases.CreateKategori(&kategori)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]interface{}{
			"code":    422,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    201,
		"status":  "success",
		"message": "success create new category",
		"data": map[string]interface{}{
			"category": newKategori,
		},
	})
}

func UpdateKategoriController(c echo.Context) error {
	kategoriId := c.Param("kategoriId")
	UpdatedAt := time.Now()

	var kategori models.Kategori

	kategori = models.Kategori{
		ID:        kategoriId,
		UpdatedAt: UpdatedAt,
	}

	c.Bind(&kategori)

	updatedKategori, err := databases.UpdateKategori(kategoriId, &kategori)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]interface{}{
			"code":    422,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success update category",
		"data": map[string]interface{}{
			"category": updatedKategori,
		},
	})
}

func DeleteKategoriController(c echo.Context) error {
	kategoriId := c.Param("kategoriId")

	_, err := databases.DeleteKategori(kategoriId)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]interface{}{
			"code":    422,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success delete category",
		"data":    nil,
	})
}

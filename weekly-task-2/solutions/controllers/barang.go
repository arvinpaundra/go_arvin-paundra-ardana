package controllers

import (
	"net/http"
	"time"

	"github.com/arvinpaundra/weekly-task-2/databases"
	"github.com/arvinpaundra/weekly-task-2/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetBarangsController(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	barangs, err := databases.GetBarangs(keyword)

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
		"message": "success get all barangs",
		"data": map[string]interface{}{
			"barang": barangs,
		},
	})
}

func GetBarangController(c echo.Context) error {
	barangId := c.Param("id")

	barang, err := databases.GetBarang(barangId)

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
		"message": "success get barang by id",
		"data": map[string]interface{}{
			"barang": barang,
		},
	})
}

func GetBarangByCategoryController(c echo.Context) error {
	kategoriId := c.Param("category_id")

	barangs, err := databases.GetBarangByCategory(kategoriId)

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
		"message": "success get all barangs",
		"data": map[string]interface{}{
			"barang": barangs,
		},
	})
}

func CreateBarangController(c echo.Context) error {
	ID := uuid.NewString()
	CreatedAt := time.Now()
	updatedAt := time.Now()

	barang := models.Barang{
		ID:        ID,
		CreatedAt: CreatedAt,
		UpdatedAt: updatedAt,
	}

	c.Bind(&barang)

	newBarang, err := databases.CreateBarang(&barang)

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
		"message": "success get all barangs",
		"data": map[string]interface{}{
			"barang": newBarang,
		},
	})
}

func UpdateBarangController(c echo.Context) error {
	barangId := c.Param("id")
	updatedAt := time.Now()

	barang := models.Barang{
		ID:        barangId,
		UpdatedAt: updatedAt,
	}

	c.Bind(&barang)

	updatedBarang, err := databases.UpdateBarang(barangId, &barang)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]interface{}{
			"code":    422,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success update barang",
		"data": map[string]interface{}{
			"barang": updatedBarang,
		},
	})
}

func DeleteBarangController(c echo.Context) error {
	barangId := c.Param("id")

	_, err := databases.DeleteBarang(barangId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success delete barang",
		"data":    nil,
	})
}

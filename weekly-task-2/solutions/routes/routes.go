package routes

import (
	"github.com/arvinpaundra/weekly-task-2/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	category := e.Group("/categories")

	category.POST("", controllers.CreatekategoriController)
	category.GET("", controllers.GetKategorisController)
	category.GET("/:kategoriId", controllers.GetKategoriController)
	category.PUT("/:kategoriId", controllers.UpdateKategoriController)
	category.DELETE("/:kategoriId", controllers.DeleteKategoriController)

	item := e.Group("/items")

	item.POST("", controllers.CreateBarangController)
	item.GET("", controllers.GetBarangsController)
	item.GET("/category/:category_id", controllers.GetBarangByCategoryController)
	item.GET("/:id", controllers.GetBarangController)
	item.PUT("/:id", controllers.UpdateBarangController)
	item.DELETE("/:id", controllers.DeleteBarangController)

	return e
}

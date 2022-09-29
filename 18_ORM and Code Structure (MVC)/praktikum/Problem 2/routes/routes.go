package routes

import (
	"github.com/arvinpaundra/rest-orm/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// users
	u := e.Group("/users")
	u.GET("", controllers.GetUsersController)
	u.POST("", controllers.CreateUserController)
	u.GET("/:id", controllers.GetUserController)
	u.PUT("/:id", controllers.UpdateUserController)
	u.DELETE("/:id", controllers.DeleteUserController)

	// books
	b := e.Group("/books")
	b.GET("", controllers.GetBooksController)
	b.POST("", controllers.CreateBookController)
	b.GET("/:id", controllers.GetBookController)
	b.PUT("/:id", controllers.UpdateBookController)
	b.DELETE("/:id", controllers.DeleteBookController)

	return e
}

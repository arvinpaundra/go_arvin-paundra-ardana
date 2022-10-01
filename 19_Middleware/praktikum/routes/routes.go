package routes

import (
	"github.com/arvinpaundra/rest-orm/constants"
	"github.com/arvinpaundra/rest-orm/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUserController)

	// users
	u := e.Group("/users")
	u.GET("", controllers.GetUsersController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.POST("", controllers.CreateUserController)
	u.GET("/:id", controllers.GetUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.PUT("/:id", controllers.UpdateUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.DELETE("/:id", controllers.DeleteUserController, middleware.JWT([]byte(constants.SECRET_JWT)))

	// books
	b := e.Group("/books")
	b.GET("", controllers.GetBooksController)
	b.POST("", controllers.CreateBookController, middleware.JWT([]byte(constants.SECRET_JWT)))
	b.GET("/:id", controllers.GetBookController)
	b.PUT("/:id", controllers.UpdateBookController, middleware.JWT([]byte(constants.SECRET_JWT)))
	b.DELETE("/:id", controllers.DeleteBookController, middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}

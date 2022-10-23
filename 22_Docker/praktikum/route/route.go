package route

import (
	"github.com/arvinpaundra/clean-architecture/config"
	"github.com/arvinpaundra/clean-architecture/controller"
	"github.com/arvinpaundra/clean-architecture/repository"
	"github.com/arvinpaundra/clean-architecture/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	e.POST("/login", userController.HandlerLogin)

	u := e.Group("/users", middleware.JWT([]byte(config.Cfg.JWTSecret)))
	u.GET("", userController.HandlerGetAllUsers)
	u.POST("", userController.HandlerCreateUser)
}

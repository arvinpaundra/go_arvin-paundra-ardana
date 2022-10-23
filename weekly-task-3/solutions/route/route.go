package route

import (
	"github.com/arvinpaundra/golang-blog/config"
	"github.com/arvinpaundra/golang-blog/controller"
	"github.com/arvinpaundra/golang-blog/repository"
	"github.com/arvinpaundra/golang-blog/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	v1 := e.Group("/api/v1")

	// user
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	v1.POST("/signup", userController.HandlerRegister)
	v1.POST("/login", userController.HandlerLogin)

	//	category
	categoryRepository := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUsecase)

	c := v1.Group("/categories")
	c.POST("", categoryController.HandlerAddCategory)
	c.GET("", categoryController.HandlerFindAll)

	//	blog
	blogRepository := repository.NewBlogRepository(db)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	blogController := controller.NewBlogController(blogUsecase)

	b := v1.Group("/blogs")
	b.GET("", blogController.HandlerFindAll)
	b.POST("", blogController.HandlerAddBlog, middleware.JWT([]byte(config.Cfg.JWTSecret)))
	b.GET("/category/:categoryId", blogController.HandlerFindByCategory)
	b.GET("/:blogId", blogController.HandlerFindById)
	b.PUT("/:blogId", blogController.HandlerUpdateBlog, middleware.JWT([]byte(config.Cfg.JWTSecret)))
	b.DELETE("/:blogId", blogController.HandlerDeleteBlog, middleware.JWT([]byte(config.Cfg.JWTSecret)))
}

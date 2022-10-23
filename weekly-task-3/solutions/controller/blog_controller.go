package controller

import (
	"github.com/arvinpaundra/golang-blog/dto"
	"github.com/arvinpaundra/golang-blog/model"
	"github.com/arvinpaundra/golang-blog/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type blogController struct {
	blogUsecase usecase.BlogUsecase
}

func NewBlogController(blogUsecase usecase.BlogUsecase) *blogController {
	return &blogController{blogUsecase}
}

func (b *blogController) HandlerFindAll(c echo.Context) error {
	title := c.QueryParam("keyword")

	blogs, err := b.blogUsecase.FindAll(title)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "sukses get data semua blog",
		"data": map[string][]model.Blog{
			"blogs": blogs,
		},
	})
}

func (b *blogController) HandlerFindById(c echo.Context) error {
	blogId := c.Param("blogId")

	blog, err := b.blogUsecase.FindById(blogId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "error",
			"message": "blog tidak ditemukan.",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "error",
		"message": "sukses get data blog",
		"data": map[string]model.Blog{
			"blog": blog,
		},
	})
}

func (b *blogController) HandlerFindByCategory(c echo.Context) error {
	categoryId := c.Param("categoryId")

	blogs, err := b.blogUsecase.FindByCategory(categoryId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "sukses get data blog berdasarkan kategori",
		"data": map[string][]model.Blog{
			"blogs": blogs,
		},
	})
}

func (b *blogController) HandlerAddBlog(c echo.Context) error {
	blogDTO := dto.BlogDTO{}

	if err := c.Bind(&blogDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Isi data dengan lengkap.",
			"data":    nil,
		})
	}

	blog, err := b.blogUsecase.Create(blogDTO)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "Berhasil mengunggah blog",
		"data": map[string]model.Blog{
			"blog": blog,
		},
	})
}

func (b *blogController) HandlerUpdateBlog(c echo.Context) error {
	blogId := c.Param("blogId")
	blogDTO := dto.BlogDTO{}

	if err := c.Bind(&blogDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Isi data dengan lengkap.",
			"data":    nil,
		})
	}

	blog, err := b.blogUsecase.Update(blogId, blogDTO)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "sukses update data blog",
		"data": map[string]model.Blog{
			"blog": blog,
		},
	})
}

func (b *blogController) HandlerDeleteBlog(c echo.Context) error {
	blogId := c.Param("blogId")

	_, err := b.blogUsecase.Delete(blogId)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "sukses hapus data blog",
		"data":    nil,
	})
}

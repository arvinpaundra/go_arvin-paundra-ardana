package usecase

import (
	"github.com/arvinpaundra/golang-blog/dto"
	"github.com/arvinpaundra/golang-blog/model"
	"github.com/arvinpaundra/golang-blog/repository"
	"github.com/google/uuid"
	"time"
)

type BlogUsecase interface {
	FindAll(title string) ([]model.Blog, error)
	FindById(blogId string) (model.Blog, error)
	FindByCategory(categoryId string) ([]model.Blog, error)
	Create(blogDTO dto.BlogDTO) (model.Blog, error)
	Update(blogId string, blogDTO dto.BlogDTO) (model.Blog, error)
	Delete(blogId string) (uint, error)
}

type blogUsecase struct {
	blogRepository repository.BlogRepository
}

func NewBlogUsecase(blogRepo repository.BlogRepository) *blogUsecase {
	return &blogUsecase{blogRepo}
}

func (b *blogUsecase) FindAll(title string) ([]model.Blog, error) {
	blogs, err := b.blogRepository.FindAll(title)

	if err != nil {
		return []model.Blog{}, err
	}

	return blogs, nil
}

func (b *blogUsecase) FindById(blogId string) (model.Blog, error) {
	blog, err := b.blogRepository.FindById(blogId)

	if err != nil {
		return model.Blog{}, err
	}

	return blog, nil
}

func (b *blogUsecase) FindByCategory(categoryId string) ([]model.Blog, error) {
	blogs, err := b.blogRepository.FindByCategory(categoryId)

	if err != nil {
		return []model.Blog{}, err
	}

	return blogs, nil
}

func (b *blogUsecase) Create(blogDTO dto.BlogDTO) (model.Blog, error) {
	id := uuid.NewString()
	createdAt := time.Now()
	updatedAt := time.Now()

	newBlog := model.Blog{
		ID:         id,
		UserId:     blogDTO.UserId,
		CategoryId: blogDTO.CategoryId,
		Title:      blogDTO.Title,
		Content:    blogDTO.Content,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}

	err := b.blogRepository.Create(newBlog)

	if err != nil {
		return model.Blog{}, err
	}

	blog, _ := b.blogRepository.FindById(id)

	return blog, nil
}

func (b *blogUsecase) Update(blogId string, blogDTO dto.BlogDTO) (model.Blog, error) {
	updatedAt := time.Now()

	updatedBlog := model.Blog{
		ID:         blogId,
		UserId:     blogDTO.UserId,
		CategoryId: blogDTO.CategoryId,
		Title:      blogDTO.Title,
		Content:    blogDTO.Content,
		UpdatedAt:  updatedAt,
	}

	err := b.blogRepository.Update(blogId, updatedBlog)

	if err != nil {
		return model.Blog{}, err
	}

	blog, _ := b.blogRepository.FindById(blogId)

	return blog, nil
}

func (b *blogUsecase) Delete(blogId string) (uint, error) {
	rowAffected, err := b.blogRepository.Delete(blogId)

	if err != nil {
		return rowAffected, err
	}

	return rowAffected, nil
}

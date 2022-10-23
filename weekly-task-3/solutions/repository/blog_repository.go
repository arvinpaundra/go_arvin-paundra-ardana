package repository

import (
	"github.com/arvinpaundra/golang-blog/database"
	"github.com/arvinpaundra/golang-blog/model"
	"gorm.io/gorm"
)

type BlogRepository interface {
	FindAll(title string) ([]model.Blog, error)
	FindById(blogId string) (model.Blog, error)
	FindByCategory(categoryId string) ([]model.Blog, error)
	Create(blog model.Blog) error
	Update(blogId string, blog model.Blog) error
	Delete(blogId string) (uint, error)
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db}
}

func (b *blogRepository) FindAll(title string) ([]model.Blog, error) {
	var blogs []model.Blog

	err := database.DB.Model(&model.Blog{}).Where("title LIKE ?", "%"+title+"%").Preload("Category").Preload("User").Find(&blogs).Error

	if err != nil {
		return []model.Blog{}, err
	}

	return blogs, nil
}

func (b *blogRepository) FindById(blogId string) (model.Blog, error) {
	blog := model.Blog{}

	err := database.DB.Model(&model.Blog{}).Where("blogs.id = ?", blogId).Preload("Category").Preload("User").Take(&blog).Error

	if err != nil {
		return model.Blog{}, err
	}

	return blog, nil
}

func (b *blogRepository) FindByCategory(categoryId string) ([]model.Blog, error) {
	var blogs []model.Blog

	err := database.DB.Model(&model.Blog{}).Where("blogs.category_id = ?", categoryId).Preload("Category").Preload("User").Find(&blogs).Error

	if err != nil {
		return []model.Blog{}, err
	}

	return blogs, nil
}

func (b *blogRepository) Create(blog model.Blog) error {
	err := database.DB.Model(&model.Blog{}).Create(&blog).Error

	if err != nil {
		return err
	}

	return nil
}

func (b *blogRepository) Update(blogId string, blog model.Blog) error {
	err := database.DB.Model(&model.Blog{}).Where("id = ?", blogId).Updates(&blog).Error

	if err != nil {
		return err
	}

	return nil
}

func (b *blogRepository) Delete(blogId string) (uint, error) {
	err := database.DB.Model(&model.Blog{}).Where("id = ?", blogId).Delete(&model.Blog{}).Error

	if err != nil {
		return 0, err
	}

	return 1, nil
}

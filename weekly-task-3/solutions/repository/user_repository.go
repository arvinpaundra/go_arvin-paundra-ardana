package repository

import (
	"github.com/arvinpaundra/golang-blog/database"
	"github.com/arvinpaundra/golang-blog/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(email string, password string) (model.User, error)
	Create(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Login(email string, password string) (model.User, error) {
	user := model.User{}

	err := database.DB.Model(&model.User{}).Where("email = ? AND password = ?", email, password).Take(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) Create(user model.User) error {
	err := database.DB.Model(&model.User{}).Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

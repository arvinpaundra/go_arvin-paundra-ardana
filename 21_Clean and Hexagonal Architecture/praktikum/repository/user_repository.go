package repository

import (
	"github.com/arvinpaundra/clean-architecture/database"
	"github.com/arvinpaundra/clean-architecture/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	Create(user model.User) error
	Login(email string, password string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User

	err := database.DB.Model(&model.User{}).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Create(user model.User) error {
	err := database.DB.Model(&model.User{}).Save(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Login(email string, password string) (model.User, error) {
	var user model.User

	err := database.DB.Model(model.User{}).Where("email = ? AND password = ?", email, password).Take(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

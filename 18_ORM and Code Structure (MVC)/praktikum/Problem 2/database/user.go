package database

import (
	"github.com/arvinpaundra/rest-orm/config"
	"github.com/arvinpaundra/rest-orm/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Model(&models.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id uint) (interface{}, error) {
	var user models.User

	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Model(&models.User{}).Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user models.User, id uint) (interface{}, error) {
	err := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(models.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(id uint) (interface{}, error) {
	if err := config.DB.Unscoped().Model(&models.User{}).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

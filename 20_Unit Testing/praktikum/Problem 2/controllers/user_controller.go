package controllers

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/unit-test/config"
	"github.com/arvinpaundra/unit-test/middlewares"
	"github.com/arvinpaundra/unit-test/models"
	"github.com/labstack/echo/v4"
)

type DTOUser struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func LoginUserController(c echo.Context) error {
	user := DTOUser{}

	c.Bind(&user)

	loggedUser := models.User{}

	err := config.DB.Model(&models.User{}).Where("email = ? AND password = ?", user.Email, user.Password).First(&loggedUser).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, _ := middlewares.CreateToken(loggedUser.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func GetUsersController(c echo.Context) error {
	users := []DTOUser{}

	err := config.DB.Model(&models.User{}).Find(&users).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"users": users,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func GetUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user := DTOUser{}

	err := config.DB.Model(&models.User{}).Where("id = ?", userId).Take(&user).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"user": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func CreateUserController(c echo.Context) error {
	user := DTOUser{}

	c.Bind(&user)

	err := config.DB.Model(&models.User{}).Save(&models.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]DTOUser{
		"user": user,
	})
}

func UpdateUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user := DTOUser{}

	c.Bind(&user)

	err := config.DB.Model(&models.User{}).Where("id = ?", userId).Updates(models.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func DeleteUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Unscoped().Model(&models.User{}).Where("id = ?", userId).Delete(&models.User{}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/arvinpaundra/rest-orm/lib/database"
	"github.com/arvinpaundra/rest-orm/models"
	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	var user models.User

	c.Bind(&user)

	userLoggedIn, err := database.LoginUser(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    userLoggedIn,
	})
}

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"users":   users,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUser(uint(userId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"user":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User

	c.Bind(&user)

	userCreated, err := database.CreateUser(user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create user",
		"user":    userCreated,
	})
}

func UpdateUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	var user models.User

	c.Bind(&user)

	userUpdated, err := database.UpdateUser(user, uint(userId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    userUpdated,
	})
}

func DeleteUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	_, err := database.DeleteUser(uint(userId))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

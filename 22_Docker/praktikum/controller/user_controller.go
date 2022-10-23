package controller

import (
	"github.com/arvinpaundra/clean-architecture/dto"
	middle "github.com/arvinpaundra/clean-architecture/middleware"
	"github.com/arvinpaundra/clean-architecture/model"
	"github.com/arvinpaundra/clean-architecture/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) HandlerLogin(c echo.Context) error {
	var userDTO dto.UserDTO

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user, err := u.userUsecase.Login(userDTO)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	token, _ := middle.CreateToken(user.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "login success",
		"data": map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}

func (u *userController) HandlerGetAllUsers(c echo.Context) error {
	users, err := u.userUsecase.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    []model.User{},
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all users",
		"data": map[string]interface{}{
			"users": users,
		},
	})
}

func (u *userController) HandlerCreateUser(c echo.Context) error {
	var user dto.UserDTO

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	err := u.userUsecase.Create(user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success create user",
		"data": map[string]dto.UserDTO{
			"user": user,
		},
	})
}

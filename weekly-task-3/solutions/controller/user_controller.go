package controller

import (
	"github.com/arvinpaundra/golang-blog/dto"
	middle "github.com/arvinpaundra/golang-blog/helper"
	"github.com/arvinpaundra/golang-blog/model"
	"github.com/arvinpaundra/golang-blog/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{userUsecase}
}

func (u *userController) HandlerLogin(c echo.Context) error {
	userDTO := dto.LoginDTO{}

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Email dan password wajib diisi.",
			"data":    nil,
		})
	}

	user, err := u.userUsecase.Login(userDTO.Email, userDTO.Password)

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
		"message": "login sukses",
		"data": map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}

func (u *userController) HandlerRegister(c echo.Context) error {
	userDTO := dto.UserDTO{}

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Isi data secara lengkap.",
			"data":    nil,
		})
	}

	user, err := u.userUsecase.Register(userDTO)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "register success",
		"data": map[string]model.User{
			"user": user,
		},
	})
}

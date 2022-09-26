package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

// ------------------------- controller -------------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	for _, user := range users {
		if user.Id == userId {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get user by id",
				"user":    user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
		"user":    nil,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	updatedUser := User{}
	c.Bind(&updatedUser)

	for i := range users {
		if users[i].Id == userId {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user",
				"user":    users[i],
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
		"user":    nil,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	for i := range users {
		if users[i].Id == userId {
			users = append(users[:i], users[i+1:]...)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success delete user",
				"user":    nil,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
		"user":    nil,
	})
}

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	err := DB.Find(&users).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all user",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// TODO get id from request params
	userId, _ := strconv.Atoi(c.Param("id"))

	// TODO create var to store user from db
	var user User

	// TODO query from db
	err := DB.Model(&User{}).Where("id = ?", userId).Take(&user).Error

	// TODO check is user with expected id exist, if not return not found
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"user":    nil,
		})
	}

	// TODO user with expected id found, return OK
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	err := DB.Save(&user).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	// TODO get id from request params
	userId, _ := strconv.Atoi(c.Param("id"))

	err := DB.Unscoped().Model(&User{}).Where("id = ?", userId).Delete(&User{}).Error

	// TODO check is id exist in db
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"user":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user deleted",
	})
}

func UpdateUserController(c echo.Context) error {
	// TODO get id user from request param
	userId, _ := strconv.Atoi(c.Param("id"))

	// TODO create var to store data user
	var user User
	// TODO bind the request body into user var
	c.Bind(&user)

	// TODO update user with the new value from var user
	err := DB.Model(&User{}).Where("id = ?", userId).Updates(User{Name: user.Name, Email: user.Email, Password: user.Password}).Error

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"user":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func main() {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

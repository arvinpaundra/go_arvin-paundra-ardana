package config

import (
	"fmt"

	"github.com/arvinpaundra/unit-test/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

type ConfigTest struct {
	DB_USERNAME_TEST string
	DB_PASSWORD_TEST string
	DB_PORT_TEST     string
	DB_HOST_TEST     string
	DB_NAME_TEST     string
}

func InitDB() {
	config := Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_PORT:     "3306",
		DB_HOST:     "localhost",
		DB_NAME:     "golang_db",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}

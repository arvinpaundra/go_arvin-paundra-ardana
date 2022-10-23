package database

import (
	"fmt"

	"github.com/arvinpaundra/golang-blog/config"
	"github.com/arvinpaundra/golang-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	username := config.Cfg.DBUsername
	password := config.Cfg.DBPassword
	address := config.Cfg.DBAddress
	dbName := config.Cfg.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, dbName)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&model.Category{}, &model.User{}, &model.Blog{})
}

package database

import (
	"fmt"
	"github.com/arvinpaundra/clean-architecture/config"
	"github.com/arvinpaundra/clean-architecture/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	address := config.Cfg.DBAddress
	username := config.Cfg.DBUsername
	password := config.Cfg.DBPassword

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/crud_go?charset=utf8mb4&parseTime=True&loc=Local", username, password, address)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&model.User{})
}

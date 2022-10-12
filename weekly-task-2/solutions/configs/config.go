package configs

import (
	"fmt"

	"github.com/arvinpaundra/weekly-task-2/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func InitDB() {
	config := Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_NAME:     "crud_go",
		DB_HOST:     "localhost",
		DB_PORT:     "3306",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	var err error

	DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	RunMigration()
}

func RunMigration() {
	DB.AutoMigrate(&models.Kategori{}, &models.Barang{})
	DB.Model(&models.Barang{}).AddForeignKey("id_kategori", "kategoris(id)", "RESTRICT", "RESTRICT")
}

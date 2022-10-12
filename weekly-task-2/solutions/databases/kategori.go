package databases

import (
	"errors"

	"github.com/arvinpaundra/weekly-task-2/configs"
	"github.com/arvinpaundra/weekly-task-2/models"
	"github.com/jinzhu/gorm"
)

func CreateKategori(kategori *models.Kategori) (interface{}, error) {
	if err := configs.DB.Model(&models.Kategori{}).Save(&kategori).Error; err != nil {
		return nil, err
	}

	return kategori, nil
}

func GetKategoris() (interface{}, error) {
	var kategori []models.Kategori

	if err := configs.DB.Model(&models.Kategori{}).Order("name").Find(&kategori).Error; err != nil {
		return nil, err
	}

	return kategori, nil
}

func GetKategori(id string) (interface{}, error) {
	var kategori models.Kategori

	err := configs.DB.Model(&models.Kategori{}).Where("id = ?", id).Limit(1).First(&kategori).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Kategori not found")
		}

		return nil, err
	}

	return kategori, nil
}

func UpdateKategori(id string, kategori *models.Kategori) (interface{}, error) {
	err := configs.DB.Model(&models.Kategori{}).Where("id = ?", id).Updates(models.Kategori{Name: kategori.Name, UpdatedAt: kategori.UpdatedAt}).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Kategori not found")
		}

		return nil, err
	}

	return kategori, nil
}

func DeleteKategori(id string) (interface{}, error) {
	err := configs.DB.Model(&models.Kategori{}).Where("id = ?", id).Delete(models.Kategori{}).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, errors.New("Kategori not found")
		}

		return 0, configs.DB.Error
	}

	return configs.DB.RowsAffected, nil
}

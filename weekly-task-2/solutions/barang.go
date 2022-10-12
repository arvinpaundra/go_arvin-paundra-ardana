package databases

import (
	"errors"
	"strings"

	"github.com/arvinpaundra/weekly-task-2/configs"
	"github.com/arvinpaundra/weekly-task-2/models"
	"github.com/jinzhu/gorm"
)

func GetBarangs(keyword string) (interface{}, error) {
	var barangs []models.Barang

	queryName := strings.Join([]string{"%", keyword, "%"}, "")

	err := configs.DB.Model(&models.Barang{}).Where("nama LIKE ?", queryName).Order("nama").Scan(&barangs).Error

	if err != nil {
		return nil, err
	}

	for i := range barangs {
		err = configs.DB.Model(&models.Kategori{}).Where("id = ?", barangs[i].IdKategori).First(&barangs[i].Kategori).Error

		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return nil, errors.New("Category not found")
			}

			return nil, err
		}
	}

	return barangs, nil
}

func GetBarang(id string) (interface{}, error) {
	var barang models.Barang

	err := configs.DB.Model(&models.Barang{}).Where("barangs.id = ?", id).Limit(1).Find(&barang).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Barang not found")
		}

		return nil, err
	}

	err = configs.DB.Model(&models.Kategori{}).Where("id = ?", barang.IdKategori).First(&barang.Kategori).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Category not found")
		}

		return nil, err
	}

	return barang, nil
}

func GetBarangByCategory(kategoriId string) (interface{}, error) {
	var barangs []models.Barang

	err := configs.DB.Model(&models.Barang{}).Where("id_kategori = ?", kategoriId).Find(&barangs).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Category not found")
		}

		return nil, err
	}

	for i := range barangs {
		err = configs.DB.Model(&models.Kategori{}).Where("id = ?", kategoriId).First(&barangs[i].Kategori).Error

		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return nil, errors.New("Category not found")
			}

			return nil, err
		}
	}

	return barangs, nil
}

func CreateBarang(barang *models.Barang) (interface{}, error) {
	err := configs.DB.Model(&models.Barang{}).Save(&barang).Error

	if err != nil {
		return nil, err
	}

	err = configs.DB.Model(&models.Kategori{}).Where("id = ?", barang.IdKategori).First(&barang.Kategori).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Category not found")
		}

		return nil, err
	}

	return barang, nil
}

func UpdateBarang(id string, barang *models.Barang) (interface{}, error) {
	err := configs.DB.Model(&models.Barang{}).Where("id = ?", id).Updates(models.Barang{Nama: barang.Nama, Deskripsi: barang.Deskripsi, JumlahStok: barang.JumlahStok, Harga: barang.Harga, IdKategori: barang.IdKategori}).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Barang not found")
		}

		return nil, err
	}

	err = configs.DB.Model(&models.Kategori{}).Where("id = ?", barang.IdKategori).First(&barang.Kategori).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("Category not found")
		}

		return nil, err
	}

	return &barang, nil
}

func DeleteBarang(id string) (interface{}, error) {
	err := configs.DB.Model(&models.Barang{}).Where("id = ?", id).Delete(models.Barang{}).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, errors.New("Barang not found")
		}

		return 0, err
	}

	return configs.DB.RowsAffected, nil
}

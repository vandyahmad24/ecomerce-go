package database

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
)

var provinces []models.Provinces
var getProvince models.Provinces

func GetProvinces() (interface{}, error) {
	if e := config.DB.Find(&provinces).Error; e != nil {
		return nil, e
	}
	return provinces, nil
}

func GetProvincesById(provinces *models.Provinces, id int) (interface{}, error) {
	if e := config.DB.First(&provinces, id).Error; e != nil {
		return nil, e
	}

	return provinces, nil
}

func DeleteProvincesById(id int) (interface{}, error) {

	if e := config.DB.Delete(&provinces, id).Error; e != nil {
		return nil, e
	}
	return "deleted", nil

}

func StoreProvinces(provinces *models.Provinces) (interface{}, error) {
	if err := config.DB.Create(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil

}

func PutProvincesByID(provinces *models.Provinces, id int, name string) (interface{}, error) {

	// config.DB.First(&provinces, id).Update({
	e := config.DB.Where("id=?", id).First(&provinces).Error
	if e != nil {
		return nil, e
	}
	// fmt.Println(e)
	provinces.Name = name
	config.DB.Save(&provinces)
	return provinces, nil
}

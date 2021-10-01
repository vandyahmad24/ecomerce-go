package database

import (
	"errors"
	"fmt"
	"go-ecommerce/config"
	"go-ecommerce/models"
)

var city []models.City

func GetCityByProvincesId(prov_id int) (interface{}, error) {

	var result []models.CityResponse
	var count int64
	e := config.DB.Table("cities").Select("cities.id as id,cities.name as city_name, cities.provinces_id, provinces.name as provinces_name").Joins("join provinces on provinces.id = cities.provinces_id").Where("cities.provinces_id = ?", prov_id).Find(&result).Count(&count).Error
	if e != nil {
		return nil, e
	}
	if count <= 0 {
		return nil, errors.New("record not found")
	}

	fmt.Println(count)
	return result, nil
}

func GetCityById(id int) (interface{}, error) {
	var result models.CityResponse
	e := config.DB.Table("cities").Select("cities.id as id,cities.name as city_name, cities.provinces_id, provinces.name as provinces_name").Joins("join provinces on provinces.id = cities.provinces_id").Where("cities.id = ?", id).Find(&result).Error
	if e != nil {
		return nil, e
	}

	return result, nil
}

func Storecity(city *models.City) (interface{}, error) {
	if err := config.DB.Create(&city).Error; err != nil {
		return nil, err
	}
	return city, nil

}

func DeleteCityById(id int) (interface{}, error) {
	if e := config.DB.Delete(&city, id).Error; e != nil {
		return nil, e
	}
	return "deleted", nil
}

func PutCityById(city *models.City, id int, name string, prov_id uint) (interface{}, error) {
	e := config.DB.Where("id=?", id).First(&city).Error
	if e != nil {
		return nil, e
	}
	// fmt.Println(e)
	city.Name = name
	city.ProvincesId = prov_id
	config.DB.Save(&city)
	return city, nil
}

func AccessGetCityById(city *models.City, id int) (*models.City, error) {
	// err := repo.db.Preload("Categories").Where("id = ?", id).First(&product).Error
	//
	if e := config.DB.Preload("Provinces").First(&city, id).Error; e != nil {
		return nil, e
	}

	return city, nil
}

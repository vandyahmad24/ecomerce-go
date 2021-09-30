package database

import (
	"errors"
	"go-ecommerce/config"
	"go-ecommerce/models"
	"strconv"
)

var city []models.City

func GetCityByProvincesId(prov_id int) (interface{}, error) {
	e := config.DB.Preload("Provinces").Where("provinces_id=?", prov_id).Find(&city).Error

	jumlah := len(city)
	if jumlah <= 0 {
		return nil, errors.New("not found city with provinces id " + strconv.Itoa(prov_id))
	}
	if e != nil {
		return nil, e
	}
	return city, nil
}

func GetCityById(city *models.City, id int) (interface{}, error) {
	// err := repo.db.Preload("Categories").Where("id = ?", id).First(&product).Error
	//
	if e := config.DB.Preload("Provinces").First(&city, id).Error; e != nil {
		return nil, e
	}

	return city, nil
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

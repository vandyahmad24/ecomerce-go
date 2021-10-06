package database

import (
	"errors"
	"go-ecommerce/config"
	"go-ecommerce/models"
)

var category []models.Categories

func GetAllCategoryActive() (interface{}, error) {
	// status := false
	e := config.DB.Table("categories").Where("status_categories = ?", true).Find(&category).Error

	jumlah := len(category)
	if jumlah <= 0 {
		return nil, errors.New("not found category")
	}
	if e != nil {
		return nil, e
	}
	return category, nil
}
func GetAllCategoryInActive() (interface{}, error) {
	// status := false
	e := config.DB.Table("categories").Where("status_categories = ?", false).Find(&category).Error

	jumlah := len(category)
	if jumlah <= 0 {
		return nil, errors.New("not found category")
	}
	if e != nil {
		return nil, e
	}
	return category, nil
}

func StoreCategory(category *models.Categories) (interface{}, error) {
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil

}

func GetCategoryById(category *models.Categories, id int) (interface{}, error) {
	if e := config.DB.First(&category, id).Error; e != nil {
		return nil, e
	}

	return category, nil
}

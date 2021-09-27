package database

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users
	if e := config.DB.Find(&users).Error; e != nil{
		return nil, e
	}
	return  users, nil
}

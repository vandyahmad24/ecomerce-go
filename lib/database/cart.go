package database

import (
	"errors"
	"fmt"
	"go-ecommerce/config"
	"go-ecommerce/models"
)

func GetListCart(id int) (interface{}, error) {
	var result []models.CartResponse
	var count int64
	e := config.DB.Table("carts").Select("carts.id, carts.products_id, carts.users_id, carts.qty, users.name, products.name as products_name, users.name as users_name").Joins("join products on products.id = carts.products_id").Joins("join users on users.id = carts.users_id").Where("carts.users_id = ?", id).Find(&result).Count(&count).Error
	// .
	if e != nil {
		return nil, e
	}
	if count <= 0 {
		return nil, errors.New("record not found")
	}

	fmt.Println(count)
	return result, nil

}

func DeleteCartById(id int) (interface{}, error) {
	var cart models.Carts
	if e := config.DB.Delete(&cart, id).Error; e != nil {
		return nil, e
	}
	return "deleted", nil

}

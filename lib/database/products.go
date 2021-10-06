package database

import (
	"errors"
	"fmt"
	"go-ecommerce/config"
	"go-ecommerce/models"
)

func StoreProduct(product *models.Products) (interface{}, error) {
	var responseProduk models.ProductsResponse
	var categori models.Categories
	if err := config.DB.Create(&product).Scan(&responseProduk).Error; err != nil {
		return nil, err
	}

	config.DB.First(&categori, responseProduk.CategoriesId)

	responseProduk.CategoriesName = categori.Name
	return responseProduk, nil

}

func GetProductByCategory(id int) (interface{}, error) {
	var result []models.ProductsResponse
	var count int64
	e := config.DB.Table("products").Select("products.name, products.id, products.slug, products.seller_id, products.price, products.description, products.categories_id, categories.name as categories_name, users.name as seller_name").Joins("join categories on categories.id = products.categories_id").Where("products.categories_id = ?", id).Joins("join users on users.id = products.seller_id").Find(&result).Count(&count).Error
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

func GetProductById(id int) (interface{}, error) {
	var result models.ProductsResponse
	var count int64
	e := config.DB.Table("products").Select("products.name, products.id, products.slug, products.seller_id, products.price, products.description, products.categories_id, categories.name as categories_name, users.name as seller_name").Joins("join categories on categories.id = products.categories_id").Where("products.id = ?", id).Joins("join users on users.id = products.seller_id").First(&result).Count(&count).Error
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

func AddCartByProdukID(cart *models.Carts, id int) (interface{}, error) {
	var cartResponse models.CartResponse
	if err := config.DB.Create(&cart).Scan(&cartResponse).Error; err != nil {
		return nil, err
	}
	return cartResponse, nil
}

package models

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID           uint
	SellerId     uint   `json:"seller_id" form:"seller_id" validate:"required"`
	Users        Users  `gorm:"foreignKey:SellerId"`
	Name         string `json:"name" form:"name" gorm:"size:255" validate:"required"`
	Price        int    `json:"price" form:"price" gorm:"size:255" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	Slug         string `json:"slug" form:"slug" gorm:"size:255"`
	CategoriesId uint   `json:"categories_id" form:"categories_id" validate:"required"`
	Categories   Categories
}

type ProductsRequest struct {
	ID           uint
	SellerId     uint   `json:"seller_id" form:"seller_id"`
	Name         string `json:"name" form:"name" validate:"required"`
	Price        int    `json:"price" form:"price" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	Slug         string `json:"slug" form:"slug"`
	CategoriesId uint   `json:"categories_id" form:"categories_id" validate:"required"`
}

type ProductsResponse struct {
	ID             uint   `json:"id" form:"id"`
	SellerId       uint   `json:"seller_id" form:"seller_id"`
	SellerName     string `json:"seller_name" form:"seller_name"`
	Name           string `json:"name" form:"name" `
	Price          int    `json:"price" form:"price" `
	Description    string `json:"description" form:"description" `
	Slug           string `json:"slug" form:"slug"`
	CategoriesId   uint   `json:"categories_id" form:"categories_id" `
	CategoriesName string `json:"categories_name"`
}

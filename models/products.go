package models


import (
	"gorm.io/gorm"
	_ "github.com/go-playground/validator/v10"
)

type Products struct {
	gorm.Model
	ID uint
	SellerId uint `json:"seller_id" form:"seller_id" validate:"required"`
	Users      Users `gorm:"foreignKey:SellerId"`
	Name string `json:"name" form:"name" gorm:"size:255" validate:"required"`
	Price int `json:"price" form:"price" gorm:"size:255" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Slug string `json:"slug" form:"slug" gorm:"size:255"`
	CategoriesId uint `json:"categories_id" form:"categories_id" validate:"required"`
	Categories Categories


}

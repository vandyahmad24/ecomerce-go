package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID uint
	SellerId uint `json:"seller_id" form:"seller_id"`
	Users      Users `gorm:"foreignKey:SellerId"`
	Name string `json:"name" form:"name" gorm:"size:255"`
	Price int `json:"price" form:"price" gorm:"size:255"`
	Description string `json:"description" form:"description"`
	Slug string `json:"slug" form:"slug" gorm:"size:255"`
	CategoriesId uint `json:"categories_id" form:"categories_id"`
	Categories Categories


}

package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	ID uint
	ProductsId uint `json:"products_id" form:"products_id"`
	Products      Products `gorm:"foreignKey:ProductsId"`
	UsersId uint `json:"users_id" form:"users_id"`
	Users      Users `gorm:"foreignKey:UsersId"`
	Qty string `json:"qty" form:"qty"`

}
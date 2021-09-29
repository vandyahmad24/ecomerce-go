package models

import (
	"gorm.io/gorm"
	_ "github.com/go-playground/validator/v10"
)

type Carts struct {
	gorm.Model
	ID uint
	ProductsId uint `json:"products_id" form:"products_id" validate:"required"`
	Products      Products `gorm:"foreignKey:ProductsId"`
	UsersId uint `json:"users_id" form:"users_id" validate:"required"`
	Users      Users `gorm:"foreignKey:UsersId"`
	Qty string `json:"qty" form:"qty" validate:"required"`

}
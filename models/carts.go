package models

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Carts struct {
	gorm.Model
	ID         uint
	ProductsId uint     `json:"products_id" form:"products_id" validate:"required"`
	Products   Products `gorm:"foreignKey:ProductsId"`
	UsersId    uint     `json:"users_id" form:"users_id" validate:"required"`
	Users      Users    `gorm:"foreignKey:UsersId"`
	Qty        string   `json:"qty" form:"qty" validate:"required"`
}
type CartsRequest struct {
	ID         uint
	ProductsId uint   `json:"products_id" form:"products_id"`
	UsersId    uint   `json:"users_id" form:"users_id"`
	Qty        string `json:"qty" form:"qty"`
}

type CartResponse struct {
	ID           uint   `json:"id" form:"id"`
	ProductsId   uint   `json:"products_id" form:"products_id"`
	ProductsName string `json:"products_name" form:"products_name"`
	UsersId      uint   `json:"users_id" form:"users_id"`
	UsersName    string `json:"users_name" form:"users_name"`
	Qty          string `json:"qty" form:"qty"`
}

type CartCheckout struct {
	ID         uint `json:"id" form:"id"`
	ProductsId uint `json:"products_id" form:"products_id"`
	Price      int  `json:"price" form:"price"`
	Qty        int  `json:"qty" form:"qty"`
}

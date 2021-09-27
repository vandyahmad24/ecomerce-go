package models


import (
	"gorm.io/gorm"
	_ "github.com/go-playground/validator/v10"
)
type TransactionDetails struct {
	gorm.Model
	ID uint
	TransactionId uint `json:"transaction_id" form:"transaction_id" validate:"required"`
	Transactions      Transactions `gorm:"foreignKey:TransactionId"`
	ProductsId uint `json:"products_id" form:"products_id" validate:"required"`
	Products      Products `gorm:"foreignKey:ProductsId"`
	Qty string `json:"qty" form:"qty" validate:"required"`
}
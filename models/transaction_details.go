package models

import "gorm.io/gorm"

type TransactionDetails struct {
	gorm.Model
	ID uint
	TransactionId uint `json:"transaction_id" form:"transaction_id"`
	Transactions      Transactions `gorm:"foreignKey:TransactionId"`
	ProductsId uint `json:"products_id" form:"products_id"`
	Products      Products `gorm:"foreignKey:ProductsId"`
	Qty string `json:"qty" form:"qty"`
}
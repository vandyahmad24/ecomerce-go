package models

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	ID uint
	UserId uint `json:"user_id" form:"user_id"`
	Users      Users `gorm:"foreignKey:UserId"`
	ShippingPrice int `json:"shipping_price" form:"shipping_price"`
	TotalPrice int `json:"total_price" form:"total_price"`
	TransactionsStatus string `json:"transactions_status" form:"transactions_status" gorm:"size:255"`
	Resi string `json:"resi" form:"resi" gorm:"size:255"`
	PaymentId uint `json:"payment_id" form:"payment_id"`
	PaymentMethods      PaymentMethods `gorm:"foreignKey:PaymentId"`
}
package models



import (
	"gorm.io/gorm"
	_ "github.com/go-playground/validator/v10"
)

type Transactions struct {
	gorm.Model
	ID uint
	UserId uint `json:"user_id" form:"user_id" validate:"required"`
	Users      Users `gorm:"foreignKey:UserId"`
	ShippingPrice int `json:"shipping_price" form:"shipping_price" validate:"required"`
	TotalPrice int `json:"total_price" form:"total_price" validate:"required"`
	TransactionsStatus string `json:"transactions_status" form:"transactions_status" gorm:"size:255" validate:"required"`
	Resi string `json:"resi" form:"resi" gorm:"size:255" validate:"required"`
	PaymentId uint `json:"payment_id" form:"payment_id" validate:"required"`
	PaymentMethods      PaymentMethods `gorm:"foreignKey:PaymentId"`
}
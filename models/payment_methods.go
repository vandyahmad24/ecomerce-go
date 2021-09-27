package models


import (
	"gorm.io/gorm"
	_ "github.com/go-playground/validator/v10"
)


type PaymentMethods struct {
	gorm.Model
	ID uint
	Name string `json:"name" form:"name" validate:"required"`
	Status bool `json:"status" form:"status" gorm:"default:0"`
}
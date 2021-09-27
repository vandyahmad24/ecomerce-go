package models

import "gorm.io/gorm"

type PaymentMethods struct {
	gorm.Model
	ID uint
	Name string `json:"name" form:"name"`
	Status bool `json:"status" form:"status" gorm:"default:0"`
}
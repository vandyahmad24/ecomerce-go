package models


import (
	"gorm.io/gorm"
	_ "github.com/go-playground/validator/v10"
)

type Provinces struct {
	gorm.Model
	ID uint
	Name string `json:"name" form:"name" gorm:"size:255" validate:"required"`
}

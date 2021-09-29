package models

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Provinces struct {
	gorm.Model
	ID   uint
	Name string `json:"name" form:"name" gorm:"size:255" validate:"required"`
}

type ProvinesId struct {
	gorm.Model
	ID uint
}

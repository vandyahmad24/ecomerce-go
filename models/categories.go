package models

import (
	"gorm.io/gorm"
	_ "github.com/go-playground/validator/v10"
)

type Categories struct {
	gorm.Model
	ID uint
	Name string `json:"name" form:"name" validate:"required"`
	Slug string `json:"slug" form:"slug" validate:"required"`
	StatusCategories bool `json:"status_categories" form:"status_categories" gorm:"default:1"`
}

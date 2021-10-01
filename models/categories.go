package models

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	ID               uint
	Name             string `json:"name" form:"name" validate:"required"`
	Slug             string `json:"slug" form:"slug" validate:"required"`
	StatusCategories bool   `json:"status_categories" form:"status_categories" `
}

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required" query:"name"`
}

package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	ID uint
	Name string `json:"name" form:"name"`
	Slug string `json:"slug" form:"slug"`
	StatusCategories bool `json:"status_categories" form:"status_categories" gorm:"default:1"`
}

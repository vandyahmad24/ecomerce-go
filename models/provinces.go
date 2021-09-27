package models

import "gorm.io/gorm"

type Provinces struct {
	gorm.Model
	ID uint
	Name string `json:"name" form:"name" gorm:"size:255"`
}

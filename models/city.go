package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	ID uint
	Name string `json:"name" form:"name" gorm:"size:255"`
	ProvincesId uint `json:"provinces_id" form:"provinces_id"`
	Provinces   Provinces


}

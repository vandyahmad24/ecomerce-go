package models

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	ID          uint
	Name        string    `json:"name" form:"name" gorm:"size:255" validate:"required"`
	ProvincesId uint      `json:"provinces_id" form:"provinces_id" validate:"required"`
	Provinces   Provinces `gorm:"foreignKey:ProvincesId" json:"provinces" `
}

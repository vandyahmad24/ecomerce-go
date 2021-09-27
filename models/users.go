package models

import (
	sql "database/sql"
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)


type (Users struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name" form:"name" gorm:"size:255" validate:"required"`
	Email string `json:"email" form:"email" gorm:"size:255" validate:"required,email"`
	Password string `json:"password" form:"password" gorm:"size:255" validate:"required"`
	ProvincesId uint `json:"provinces_id" form:"provinces_id" validate:"required"`
	Provinces   Provinces
	CityId uint `json:"city_id" form:"city_id" validate:"required"`
	City   City
	ZipCode uint `json:"zip_code" form:"zip_code" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	StoreName sql.NullString `json:"store_name" form:"store_name" gorm:"size:255"`
	StoreStatus bool `json:"store_status" form:"store_status" gorm:"default:0"`
	OtpCode sql.NullString `json:"otp_code" form:"otp_code" gorm:"size:14"`
	IsActive bool `json:"is_active" form:"is_active" gorm:"default:0"`
	IsAdmin bool `json:"is_admin" form:"is_admin" gorm:"default:0"`
	Token string `json:"token" form:"token"`
}


)



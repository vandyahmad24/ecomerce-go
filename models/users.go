package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `json:"name" form:"name" gorm:"size:255"`
	Email string `json:"email" form:"email" gorm:"size:255"`
	Password string `json:"password" form:"password" gorm:"size:255"`
	ProvincesId uint `json:"provinces_id" form:"provinces_id"`
	Provinces   Provinces
	CityId uint `json:"city_id" form:"city_id"`
	City   City
	ZipCode uint `json:"zip_code" form:"zip_code"`
	Address string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	StoreName string `json:"store_name" form:"store_name" gorm:"size:255"`
	StoreStatus bool `json:"store_status" form:"store_status" gorm:"default:0"`
	OtpCode uint `json:"otp_code" form:"otp_code" gorm:"size:14"`
	IsActive bool `json:"is_active" form:"is_active" gorm:"default:0"`
	IsAdmin bool `json:"is_admin" form:"is_admin" gorm:"default:0"`


}
